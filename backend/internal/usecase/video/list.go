package video

import (
	"context"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

func (uc *VideoUseCase) ListMine(
	ctx context.Context,
	ownerID uuid.UUID,
	limit int,
) ([]*video_domain.Video, error) {
	// TODO: 実装
	return nil, nil
}

func (uc *VideoUseCase) ListPublic(
	ctx context.Context,
	limit int,
) ([]*video_domain.Video, error) {
	// TODO: 実装
	return nil, nil
}

func (uc *VideoUseCase) SearchByTag(
	ctx context.Context,
	tag video_value.Tag,
	limit int,
) ([]*video_domain.Video, error) {
	// TODO: 実装
	return nil, nil
}
