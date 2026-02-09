package video

import (
	"context"
	"fmt"
	"io"

	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

func (uc *VideoUseCase) UploadSource(
	ctx context.Context,
	videoID uuid.UUID,
	videoData io.Reader,
) error {

	video, err := uc.videoRepo.FindByID(ctx, videoID)
	if err != nil {
		return err
	}

	// 冪等性確保
	if video.Status() != video_value.StatusInitial {
		return nil
	}

	sourceKey := fmt.Sprintf(
		"videos/%s/source",
		videoID.String(),
	)

	if err := uc.storage.SaveSource(ctx, sourceKey, videoData); err != nil {
		return err
	}

	if err := video.MarkUploaded(sourceKey); err != nil {
		// NOTE: best-effort cleanup. orphaned data may remain.
		_ = uc.storage.Delete(ctx, sourceKey)
		return err
	}

	if err := uc.videoRepo.Save(ctx, video); err != nil {
		// NOTE: best-effort cleanup. orphaned data may remain.
		_ = uc.storage.Delete(ctx, sourceKey)
		return err
	}

	return nil
}
