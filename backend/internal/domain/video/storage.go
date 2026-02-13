package video

import (
	"context"
	"io"
	"time"
)

type Storage interface {
	// 元動画の保存（外部アップロード）
	SaveSource(ctx context.Context, sourceKey string, data io.Reader) error

	// 変換後動画の保存（Transcodeプロセスからのアップロード）
	SaveStream(ctx context.Context, streamKey string, data io.Reader) error
	// 一時アクセスURLの取得
	GenerateTemporaryAccessURL(ctx context.Context, streamKey string, expires time.Duration) (string, error)
	GetStream(ctx context.Context, streamKey string) (io.ReadSeeker, error)
	// 元動画の削除
	DeleteSource(ctx context.Context, sourceKey string) error
	// 変換後動画の削除
	DeleteStream(ctx context.Context, streamKey string) error
}
