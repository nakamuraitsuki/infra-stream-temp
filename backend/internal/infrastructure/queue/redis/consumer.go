package redis

import (
	"context"
	"log"
	"sync"

	"example.com/m/internal/usecase/job"
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

func NewConsumer(client *Client, queue job.Queue, key string) job.Consumer {
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
