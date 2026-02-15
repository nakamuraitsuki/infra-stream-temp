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

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
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

			// helper関数に処理を任せる
			// TODO: ゴルーチンリークまっしぐらなので、worker pool化するなどの対策が必要
			go c.dispatch(ctx, []byte(res[1]))
		}
	}
}

func (c *consumer) dispatch(ctx context.Context, data []byte) {
	var msg jobMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Printf("Failed to unmarshal job message: %v", err)
		return
	}

	c.mu.RLock()
	handler, ok := c.handlers[msg.Meta.Type]
	c.mu.RUnlock()

	if !ok {
		log.Printf("No handler registered for job type: %s", msg.Meta.Type)
		return
	}

	if err := handler.Handle(ctx, msg.Meta, msg.Payload); err != nil {
		log.Printf("Error handling job [type: %s, id: %s]: %v", msg.Meta.Type, msg.Meta.ID, err)
		// NOTE: retry
		if msg.Meta.Attempt < msg.Meta.MaxRetry {
			msg.Meta.Attempt++
			if err := c.queue.Enqueue(ctx, msg.Meta, msg.Payload); err != nil {
				log.Printf("Failed to re-enqueue job [type: %s, id: %s]: %v", msg.Meta.Type, msg.Meta.ID, err)
			}
		}
	}
}
