package video

import (
	"context"
	"fmt"

	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

func (uc *VideoUseCase) StartTranscoding(
	ctx context.Context,
	videoID uuid.UUID,
) error {

	video, err := uc.videoRepo.FindByID(ctx, videoID)
	if err != nil {
		return err
	}

	// 冪等性確保
	if video.Status() != video_value.StatusUploaded {
		return nil
	}

	streamKey := fmt.Sprintf(
		"videos/%s/stream",
		videoID.String(),
	)

	if err := video.StartTranscoding(streamKey); err != nil {
		return err
	}

	if err := uc.videoRepo.Save(ctx, video); err != nil {
		return err
	}

	return uc.transcoder.Transcode(
		ctx,
		video.SourceKey(),
		streamKey,
	)
}
