package redis

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"example.com/m/internal/usecase/job"
	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/errgroup"
)

const (
	TRANSCODE_WORKER_POOL_SIZE = 4
	JOB_CHANNEL_BUFFER_SIZE    = TRANSCODE_WORKER_POOL_SIZE * 2 // ワーカーが待機する分のバッファを確保
)

type consumer struct {
	client   *Client
	queue    job.Queue
	key      string // Queueの識別子
	handlers map[string]job.Handler
	mu       sync.RWMutex
}

func NewConsumer(client *Client, queue job.Queue, key string) job.Consumer {
	return &consumer{
		client:   client,
		queue:    queue,
		key:      key,
		handlers: make(map[string]job.Handler),
	}
}

func (c *consumer) Register(jobType string, h job.Handler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers[jobType] = h
}

func (c *consumer) Start(ctx context.Context) error {
	log.Printf("Starting Redis Consumer [key: %s]", c.key)
	jobCh := make(chan jobMessage, JOB_CHANNEL_BUFFER_SIZE)
	eg, gCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return c.workerPool(gCtx, TRANSCODE_WORKER_POOL_SIZE, jobCh)
	})

	eg.Go(func() error {
		defer close(jobCh) // watcherが終わったらworkerに終了を通知
		return c.watcher(gCtx, jobCh)
	})

	return eg.Wait()
}

func (c *consumer) watcher(
	ctx context.Context,
	jobCh chan<- jobMessage,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			res, err := c.client.Universal.BRPop(ctx, 5*time.Second, c.key).Result()
			if err != nil {
				if errors.Is(err, redis.Nil) {
					continue // timeout
				}
				log.Printf("BRPop error: %v", err)
				time.Sleep(1 * time.Second) // エラー時は少し休む
				continue
			}

			// NOTE: res[0]がkey, res[1]が実データなはずなので
			if len(res) < 2 {
				continue
			}

			var msg jobMessage
			if err := json.Unmarshal([]byte(res[1]), &msg); err != nil {
				log.Printf("Failed to unmarshal job message: %v", err)
				continue
			}

			select {
			case jobCh <- msg:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func (c *consumer) workerPool(
	ctx context.Context,
	numWorkers int,
	jobCh <-chan jobMessage,
) error {
	eg, ctx := errgroup.WithContext(ctx)

	for range numWorkers {
		eg.Go(func() error {
			for {
				select {
				// -- contextキャンセルの検知 --
				case <-ctx.Done():
					return nil

				// -- ジョブの処理 --
				case msgByte, ok := <-jobCh:
					if !ok {
						return nil // channelが閉じられた場合
					}
					if err := c.dispatch(ctx, msgByte); err != nil {
						return err
					}
				}
			}
		})
	}

	return eg.Wait()
}

func (c *consumer) dispatch(ctx context.Context, jobMsg jobMessage) error {
	c.mu.RLock()
	handler, ok := c.handlers[jobMsg.Meta.Type]
	c.mu.RUnlock()

	if !ok {
		log.Printf("No handler registered for job type: %s", jobMsg.Meta.Type)
		// ハンドラ未登録の場合でもジョブを消失させないため、元のキューに再投入する
		if err := c.queue.Enqueue(ctx, jobMsg.Meta, jobMsg.Payload); err != nil {
			log.Printf("Failed to re-enqueue job without registered handler [type: %s, id: %s]: %v", jobMsg.Meta.Type, jobMsg.Meta.ID, err)
		}
		// NOTE: リトライ処理を図るので、Consumer群に伝播させない
		return nil
	}

	if err := handler.Handle(ctx, jobMsg.Meta, jobMsg.Payload); err != nil {
		log.Printf("Error handling job [type: %s, id: %s]: %v", jobMsg.Meta.Type, jobMsg.Meta.ID, err)
		// NOTE: retry
		if jobMsg.Meta.Attempt < jobMsg.Meta.MaxRetry {
			jobMsg.Meta.Attempt++
			if err := c.queue.Enqueue(ctx, jobMsg.Meta, jobMsg.Payload); err != nil {
				log.Printf("Failed to re-enqueue job [type: %s, id: %s]: %v", jobMsg.Meta.Type, jobMsg.Meta.ID, err)
			}
		}
	}
	return nil
}
