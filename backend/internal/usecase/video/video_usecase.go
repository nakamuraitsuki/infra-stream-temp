package video

import (
	"context"
	"io"

	domain "example.com/m/internal/domain/video"
	"example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

type VideoUseCaseInterface interface {
	// 動画投稿
	Create(
		ctx context.Context,
		ownerID uuid.UUID,
		sourceKey string,
		title string,
		description string,
		tags []value.Tag,
		visibility value.Visibility,
	) (*domain.Video, error)

	// 自分の動画一覧
	ListMine(ctx context.Context, ownerID uuid.UUID, limit int) ([]*domain.Video, error)

	// 公開動画一覧（視聴用）
	ListPublic(ctx context.Context, limit int) ([]*domain.Video, error)

	// タグ検索
	SearchByTag(ctx context.Context, tag value.Tag, limit int) ([]*domain.Video, error)

	// 提供に適切なように変換
	StartTranscoding(ctx context.Context, videoID uuid.UUID, streamKey string) error

	// 動画情報取得
	GetByID(ctx context.Context, videoID uuid.UUID) (*domain.Video, error)

	// 配信
	// NOTE: 返り値のio.ReadSeekerはクライアントに閉じてもらうこと
	GetVideoStream(ctx context.Context, videoID uuid.UUID) (io.ReadSeeker, string, error)
}
