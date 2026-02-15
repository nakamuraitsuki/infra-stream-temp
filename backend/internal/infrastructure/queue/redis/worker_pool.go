package redis

import (
	"context"

	"golang.org/x/sync/errgroup"
)

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
