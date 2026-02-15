package ffmpeg

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"
)

func (t *ffmpegTranscoder) workerPool(
	ctx context.Context,
	numWorks int,
	pathCh <-chan string,
	streamKey string,
) error {

	eg, ctx := errgroup.WithContext(ctx)

	// NOTE: select構文は個人的に見づらいので
	//       caseごとにコメントを入れるようにしてほしいです。
	for i := 0; i < numWorks; i++ {
		eg.Go(func() error {
			for {
				select {
				// -- contextキャンセルの検知 --
				case <-ctx.Done():
					return ctx.Err() // contextがキャンセルされたら終了

				// -- アップロード --
				case path, ok := <-pathCh:
					if !ok {
						return nil // channelが閉じられた場合
					}
					if err := t.uploadFile(ctx, path, streamKey); err != nil {
						return err
					}
				}
			}
		})
	}

	return eg.Wait()
}

func (t *ffmpegTranscoder) uploadFile(
	ctx context.Context,
	filePath, streamKeyPrefix string,
) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileName := filepath.Base(filePath)
	s3Key := fmt.Sprintf("%s/%s", strings.TrimRight(streamKeyPrefix, "/"), fileName)

	return t.storage.SaveStream(ctx, s3Key, file)
}
