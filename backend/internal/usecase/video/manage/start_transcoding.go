package manage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	video_value "example.com/m/internal/domain/video/value"
	"example.com/m/internal/usecase/job"
	"example.com/m/internal/usecase/video/process"
	"github.com/google/uuid"
)

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

	if err := video.StartTranscoding(streamKey); err != nil {
		return err
	}

	err = uc.VideoRepo.Save(ctx, video)
	if err != nil {
		return err
	}

	payload := process.TranscodePayload{
		VideoID: videoID,
	}

	data, _ := json.Marshal(payload)

	meta := job.Metadata{
		ID:        uuid.New(),
		Type:      "video_transcode",
		Attempt:   0,
		MaxRetry:  3,
		CreatedAt: time.Now(),
	}

	err = uc.JobQueue.Enqueue(ctx, meta, data)
	if err != nil {
		// ジョブの登録に失敗したら動画の状態をアップロード済みに戻す
		if rollbackErr := video.RollbackToUploaded(); rollbackErr != nil {
			return fmt.Errorf("failed to enqueue job: %v; also failed to rollback video status: %v", err, rollbackErr)
		}
		if saveErr := uc.VideoRepo.Save(ctx, video); saveErr != nil {
			return fmt.Errorf("failed to enqueue job: %v; also failed to save rolled back video status: %v", err, saveErr)
		}
		return err
	}

	return nil
}
