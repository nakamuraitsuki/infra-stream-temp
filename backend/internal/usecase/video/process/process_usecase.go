package process

import (
	"context"

	"example.com/m/internal/domain/video"
	"github.com/google/uuid"
)

type VideoProcessUseCaseInterface interface {
	// Handle processes the video with the given ID.
	Handle(ctx context.Context, videoID uuid.UUID) error
}

type VideoProcessUseCase struct {
	VideoRepo video.Repository
	Transcoder video.Transcoder
}
