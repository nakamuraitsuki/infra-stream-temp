package process

import (
	"context"
	"log"

	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

func (uc *VideoProcessUseCase) Handle(ctx context.Context, videoID uuid.UUID, isFinalAttempt bool) error {
	video, err := uc.VideoRepo.FindByID(ctx, videoID)
	if err != nil {
		return err
	}

	if video.Status() != video_value.StatusProcessing {
		return nil
	}

	err = uc.Transcoder.Transcode(
		ctx,
		video.SourceKey(),
		video.StreamKey(),
	)
	if err != nil {
		if isFinalAttempt {
			video.MarkTranscodeFailed(video_value.FailureTranscode)
			_ = uc.VideoRepo.Save(ctx, video)
		}
		log.Println("transcode error:", err)
		return err
	}

	video.MarkTranscodeSucceeded()
	if err := uc.VideoRepo.Save(ctx, video); err != nil {
		log.Println("failed to save video after transcoding:", err)
	}
	return nil
}
