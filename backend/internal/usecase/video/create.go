package video

import (
	"context"

	video_domain "example.com/m/internal/domain/video"
	"example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

func (uc *VideoUseCase) Create(
	ctx context.Context,
	ownerID uuid.UUID,
	sourceKey string,
	title string,
	description string,
	tags []value.Tag,
	visibility value.Visibility,
) (*video_domain.Video, error) {
	// TODO: 実装
	return nil, nil
}
