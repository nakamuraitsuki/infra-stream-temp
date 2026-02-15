package redis

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

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
