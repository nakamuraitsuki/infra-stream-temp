package video

import (
	"context"
	"io"
)

type Storage interface {
	// 元動画の保存（外部アップロード）
	SaveSource(ctx context.Context, key string, data io.Reader) error
 
	// 変換後動画の保存（Transcodeプロセスからのアップロード）
	SaveStream(ctx context.Context, key string, data io.Reader) error
	GetStream(ctx context.Context, key string) (io.ReadSeeker, error)
	Delete(ctx context.Context, key string) error
}
