package manage

import (
	"context"
	"fmt"
	"io"
	"time"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"example.com/m/internal/usecase/video/query"
	"github.com/google/uuid"
)

func (uc *VideoManagementUseCase) Create(
	ctx context.Context,
	ownerID uuid.UUID,
	title string,
	description string,
	tags []video_value.Tag,
) (*video_domain.Video, error) {

	videoID := uuid.New()

	status := video_value.StatusInitial

	// NOTE: sourceKey and streamKey will be set after the video data is uploaded.
	sourceKey := ""
	streamKey := ""

	visibility := video_value.VisibilityPrivate

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

	if err := uc.VideoRepo.Save(ctx, video); err != nil {
		return nil, err
	}

	return video, nil
}

func (uc *VideoManagementUseCase) UploadSource(
	ctx context.Context,
	videoID uuid.UUID,
	videoData io.Reader,
) error {

	video, err := uc.VideoRepo.FindByID(ctx, videoID)
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

	if err := uc.Storage.SaveSource(ctx, sourceKey, videoData); err != nil {
		return err
	}

	if err := video.MarkUploaded(sourceKey); err != nil {
		// NOTE: best-effort cleanup. orphaned data may remain.
		_ = uc.Storage.Delete(ctx, sourceKey)
		return err
	}

	if err := uc.VideoRepo.Save(ctx, video); err != nil {
		// NOTE: best-effort cleanup. orphaned data may remain.
		_ = uc.Storage.Delete(ctx, sourceKey)
		return err
	}

	return nil
}

func (uc *VideoManagementUseCase) StartTranscoding(
	ctx context.Context,
	videoID uuid.UUID,
) error {

	video, err := uc.VideoRepo.FindByID(ctx, videoID)
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

	// 先にトランスコードを実行し、成功した場合のみ状態を遷移させて永続化する
	if err := uc.Transcoder.Transcode(
		ctx,
		video.SourceKey(),
		streamKey,
	); err != nil {
		return err
	}

	if err := video.StartTranscoding(streamKey); err != nil {
		return err
	}

	return uc.VideoRepo.Save(ctx, video)
}

func (uc *VideoManagementUseCase) ListMine(
	ctx context.Context,
	ownerID uuid.UUID,
	query query.VideoSearchQuery,
) ([]*video_domain.Video, error) {

	cond := video_domain.ListCondition{
		OwnerID: &ownerID,
		Limit:   query.Limit,
	}

	return uc.VideoRepo.FindByCondition(ctx, cond)
}
