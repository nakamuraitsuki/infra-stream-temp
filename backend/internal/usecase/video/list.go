package video

import (
	"context"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

// VideoSearchQuery defines the parameters for searching videos.
// example: limit, offset, sort order, etc.
type VideoSearchQuery struct {
	Limit int
	// add more fields as needed
}

func (uc *VideoUseCase) ListMine(
	ctx context.Context,
	ownerID uuid.UUID,
	query VideoSearchQuery,
) ([]*video_domain.Video, error) {
	// TODO: 実装
	return nil, nil
}

func (uc *VideoUseCase) ListPublic(
	ctx context.Context,
	query VideoSearchQuery,
) ([]*video_domain.Video, error) {
	// TODO: 実装
	return nil, nil
}

func (uc *VideoUseCase) SearchByTag(
	ctx context.Context,
	tag video_value.Tag,
	query VideoSearchQuery,
) ([]*video_domain.Video, error) {
	// TODO: 実装
	return nil, nil
}
