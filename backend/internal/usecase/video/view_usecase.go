package video

import (
	"context"
	"io"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

type VideoViewingUseCaseInterface interface {
	// GetByID returns a video by its ID.
	GetByID(ctx context.Context, videoID uuid.UUID) (*video_domain.Video, error)

	// ListPublic returns a list of publicly available videos for viewing.
	ListPublic(ctx context.Context, query VideoSearchQuery) ([]*video_domain.Video, error)

	// SearchByTag returns a list of videos matching the specified tag.
	SearchByTag(ctx context.Context, tag video_value.Tag, query VideoSearchQuery) ([]*video_domain.Video, error)

	// GetPlaybackInfo returns playback information for the specified video.
	GetPlaybackInfo(ctx context.Context, videoID uuid.UUID) (*PlaybackInfo, error)

	// GetVideoStream returns a readable video stream and its MIME type.
	GetVideoStream(ctx context.Context, videoID uuid.UUID) (io.ReadSeeker, string, error)
}

type videoViewingUseCase struct {
	videoRepo video_domain.Repository
	storage   video_domain.Storage
}
