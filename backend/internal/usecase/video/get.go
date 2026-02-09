package video

import (
	"context"

	video_domain "example.com/m/internal/domain/video"
	"github.com/google/uuid"
)

func (uc *VideoUseCase) GetByID(
	ctx context.Context,
	videoID uuid.UUID,
) (*video_domain.Video, error) {
	// TODO: 実装
	return nil, nil
}
