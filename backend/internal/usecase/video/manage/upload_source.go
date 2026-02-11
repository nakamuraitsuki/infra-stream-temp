package manage

import (
	"context"
	"fmt"
	"io"

	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

func (uc *VideoManagementUseCase) UploadSource(
	ctx context.Context,
	videoID uuid.UUID,
	videoData io.Reader,
) error {
	sourceKey := fmt.Sprintf(
		"videos/%s/source",
		videoID.String(),
	)
	
	err := uc.UoW.Do(ctx, func(ctx context.Context) error {
		video, err := uc.VideoRepo.FindByID(ctx, videoID)
		if err != nil {
			return err
		}

		if video.Status() != video_value.StatusInitial {
			return nil
		}

		if err := video.MarkUploaded(sourceKey); err != nil {
			return err
		}

		return uc.VideoRepo.Save(ctx, video)
	})

	if err != nil {
		// DBが失敗した時だけ、保存してしまったファイルを消す
		_ = uc.Storage.Delete(ctx, sourceKey)
		return err
	}

	return nil
}
