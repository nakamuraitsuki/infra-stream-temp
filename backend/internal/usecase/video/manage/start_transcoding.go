package manage

import (
	"context"
	"fmt"

	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

func (uc *VideoManagementUseCase) StartTranscoding(
	ctx context.Context,
	videoID uuid.UUID,
) error {
	return uc.UoW.Do(ctx, func(ctx context.Context) error {
		video, err := uc.VideoRepo.FindByID(ctx, videoID)
		if err != nil {
			return err
		}

		if video.Status() != video_value.StatusUploaded {
			return nil
		}

		streamKey := fmt.Sprintf("videos/%s/stream", videoID.String())
		if err := video.StartTranscoding(streamKey); err != nil {
			return err
		}

		if err := uc.VideoRepo.Save(ctx, video); err != nil {
			return err
		}

		ev := video.PullEvents()
		if err := uc.OutboxRepo.Save(ctx, ev); err != nil {
			return err
		}

		return nil
	})
}
