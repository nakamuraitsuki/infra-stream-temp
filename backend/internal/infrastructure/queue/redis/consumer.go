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
	client        *Client
	queue         job.Queue
	key           string // pending Queue の識別子
	processingKey string // processing Queue の識別子
	dlqKey        string // DLQ の識別子
	handlers      map[string]job.Handler
	mu            sync.RWMutex
}

func NewConsumer(client *Client, queue job.Queue, key string, processingKey string) job.Consumer {
	return &consumer{
		client:        client,
		queue:         queue,
		key:           key,
		processingKey: key + ":processing",
		dlqKey:        key + ":dlq",
		handlers:      make(map[string]job.Handler),
	}
}

func (c *consumer) Register(jobType string, h job.Handler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers[jobType] = h
}

func (c *consumer) Start(ctx context.Context) error {
	// NOTE: サービス再起動時などにprocessing Queueに残っているジョブをpending Queueに戻す
	for {
		err := c.client.Universal.RPopLPush(ctx, c.processingKey, c.key).Err()
		if err != nil {
			break // redis.Nil なら空なので終了
		}
	}

	log.Printf("Starting Redis Consumer [key: %s, processingKey: %s, dlqKey: %s]", c.key, c.processingKey, c.dlqKey)

	jobCh := make(chan []byte, JOB_CHANNEL_BUFFER_SIZE)
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
	jobCh chan<- []byte,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			res, err := c.client.Universal.BRPopLPush(ctx, c.key, c.processingKey, 5*time.Second).Result()
			if err != nil {
				if errors.Is(err, redis.Nil) {
					continue // timeout
				}
				log.Printf("BRPop error: %v", err)
				time.Sleep(1 * time.Second) // エラー時は少し休む
				continue
			}

			select {
			case jobCh <- []byte(res):
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func (c *consumer) workerPool(
	ctx context.Context,
	numWorkers int,
	jobCh <-chan []byte,
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
				case job, ok := <-jobCh:
					if !ok {
						return nil // channelが閉じられた場合
					}
					if err := c.dispatch(ctx, job); err != nil {
						return err
					}
				}
			}
		})
	}

	return eg.Wait()
}

func (c *consumer) dispatch(ctx context.Context, jobBytes []byte) error {
		var jobMsg jobMessage
	if err := json.Unmarshal(jobBytes, &jobMsg); err != nil {
		log.Printf("Failed to unmarshal job message: %v", err)
		return err
	}

	defer func() {
		// ジョブの処理が終わったらprocessing Queueから削除する
		if err := c.client.Universal.LRem(ctx, c.processingKey, 1, jobBytes).Err(); err != nil {
			log.Printf("Failed to remove job from processing queue [type: %s, id: %s]: %v", jobMsg.Meta.Type, jobMsg.Meta.ID, err)
		}
	}()

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

			// NOTE: 指数バックオフの計算: 2^attempt * 1秒 (例: 2s, 4s, 8s...)
			backoff := time.Duration(1<<jobMsg.Meta.Attempt) * time.Second
			log.Printf("Retrying job %s in %v (Attempt: %d)", jobMsg.Meta.ID, backoff, jobMsg.Meta.Attempt)
			
			// NOTE: 本来は「遅延キュー」に入れるのがベストだが......
			time.Sleep(backoff)

			if err := c.queue.Enqueue(ctx, jobMsg.Meta, jobMsg.Payload); err != nil {
				log.Printf("Failed to re-enqueue job [type: %s, id: %s]: %v", jobMsg.Meta.Type, jobMsg.Meta.ID, err)
			}
		} else {
			// 最大リトライ回数を超えた場合、DLQ にジョブを移動する
			if err := c.client.Universal.LPush(ctx, c.dlqKey, jobBytes).Err(); err != nil {
				log.Printf("Failed to move job to DLQ [type: %s, id: %s]: %v", jobMsg.Meta.Type, jobMsg.Meta.ID, err)
			}
		}
	}
	return nil
}


