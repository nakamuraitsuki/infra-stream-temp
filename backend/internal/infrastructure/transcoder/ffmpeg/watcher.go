package ffmpeg

import (
	"context"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func (t *ffmpegTranscoder) watchAndQueue(
	ctx context.Context,
	dir string,
	pathCh chan<- string,
) error {
	// cf. https://pkg.go.dev/github.com/fsnotify/fsnotify
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := watcher.Add(dir); err != nil {
		return err
	}

	for {
		select {
		// -- contextキャンセルの検知 --
		case <-ctx.Done():
			return ctx.Err()

		// -- ファイルイベントの検知 --
		case event, ok := <-watcher.Events:
			if !ok {
				return nil // watcherが閉じられた場合
			}

			// NOTE: ffmpegでtemp_fileフラグを使う前提
			//       Rename を検知してアップロードする
			if event.Op&fsnotify.Rename == fsnotify.Rename {
				if strings.HasSuffix(event.Name, ".ts") {
					select {
					case pathCh <- event.Name:
					case <-ctx.Done():
						return nil // contextがキャンセルされたら終了
					}
				}
			}

		// -- watcherのエラー検知 --
		case err, ok := <-watcher.Errors:
			if !ok {
				return nil // watcherが閉じられた場合
			}
			return err
		}
	}
}
