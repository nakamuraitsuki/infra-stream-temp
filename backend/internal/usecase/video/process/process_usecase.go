package process

import (
	"context"

	"example.com/m/internal/domain/video"
	"github.com/google/uuid"
)

type VideoProcessUseCaseInterface interface {
	// Handle processes the video with the given ID.
	Handle(ctx context.Context, videoID uuid.UUID, isFinalAttempt bool) error
}

type VideoProcessUseCase struct {
	VideoRepo video.Repository
	Transcoder video.Transcoder
}

func NewVideoProcessUseCase(
	videoRepo video.Repository,
	transcoder video.Transcoder,
) VideoProcessUseCaseInterface {
	return &VideoProcessUseCase{
		VideoRepo:  videoRepo,
		Transcoder: transcoder,
	}
}