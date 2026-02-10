package view

import (
	"context"
	"errors"
	"io"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"example.com/m/internal/usecase/video/query"
	"github.com/google/uuid"
)

var (
	ErrVideoNotReady  = errors.New("video is not ready for playback")
	ErrVideoForbidden = errors.New("video is not accessible")
)

type VideoViewingUseCaseInterface interface {
	// GetByID returns a video by its ID.
	GetByID(ctx context.Context, videoID uuid.UUID) (*GetByIDResult, error)

	// ListPublic returns a list of publicly available videos for viewing.
	ListPublic(ctx context.Context, query query.VideoSearchQuery) (*ListPublicResult, error)

	// SearchByTag returns a list of videos matching the specified tag.
	SearchByTag(ctx context.Context, tag video_value.Tag, query query.VideoSearchQuery) (*GetByTagsResults, error)

	// GetPlaybackInfo returns playback information for the specified video.
	GetPlaybackInfo(ctx context.Context, videoID uuid.UUID) (*PlaybackInfo, error)

	// GetVideoStream returns a readable video stream and its MIME type.
	GetVideoStream(ctx context.Context, videoID uuid.UUID) (io.ReadSeeker, string, error)
}

type VideoViewingUseCase struct {
	VideoRepo video_domain.Repository
	Storage   video_domain.Storage
}
