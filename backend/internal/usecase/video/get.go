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

	video, err := uc.videoRepo.FindByID(ctx, videoID)
	if err != nil {
		return nil, err
	}

	return video, nil
}
