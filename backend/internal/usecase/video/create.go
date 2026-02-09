package video

import (
	"context"
	"time"

	video_domain "example.com/m/internal/domain/video"
	"example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

func (uc *VideoUseCase) Create(
	ctx context.Context,
	ownerID uuid.UUID,
	title string,
	description string,
	tags []value.Tag,
) (*video_domain.Video, error) {

	videoID := uuid.New()

	status := value.StatusInitial

	// NOTE: sourceKey and streamKey will be set after the video data is uploaded.
	sourceKey := ""
	streamKey := ""

	visibility := value.VisibilityPrivate

	video := video_domain.NewVideo(
		videoID,
		ownerID,
		sourceKey,
		streamKey,
		status,
		title,
		description,
		tags,
		visibility,
		time.Now(),
	)

	if err := uc.videoRepo.Save(ctx, video); err != nil {
		return nil, err
	}

	return video, nil
}
