package video

import (
	"context"
	"io"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

type VideoUseCaseInterface interface {
	// Create uploads a new video meta data. (not the raw video data)
	Create(
		ctx context.Context,
		ownerID uuid.UUID,
		title string,
		description string,
		tags []video_value.Tag,
	) (*video_domain.Video, error)

	// UploadSource uploads the raw video data for the specified video.
	UploadSource(ctx context.Context, videoID uuid.UUID, videoData io.Reader) error

	// StartTranscoding initiates the transcoding process for a video.
	StartTranscoding(ctx context.Context, videoID uuid.UUID) error

	// ListMine returns a list of videos owned by the specified user.
	ListMine(ctx context.Context, ownerID uuid.UUID, query VideoSearchQuery) ([]*video_domain.Video, error)

	// ListPublic returns a list of publicly available videos for viewing.
	ListPublic(ctx context.Context, query VideoSearchQuery) ([]*video_domain.Video, error)

	// SearchByTag returns a list of videos matching the specified tag.
	SearchByTag(ctx context.Context, tag video_value.Tag, query VideoSearchQuery) ([]*video_domain.Video, error)

	// GetByID returns a video by its ID.
	GetByID(ctx context.Context, videoID uuid.UUID) (*video_domain.Video, error)

	// GetVideoStream returns a readable video stream and its MIME type.
	GetVideoStream(ctx context.Context, videoID uuid.UUID) (io.ReadSeeker, string, error)

	// GetPlaybackInfo returns playback information for the specified video.
	GetPlaybackInfo(ctx context.Context, videoID uuid.UUID) (*PlaybackInfo, error)
}

type VideoUseCase struct {
	videoRepo  video_domain.Repository
	storage    video_domain.Storage
	transcoder video_domain.Transcoder
}

func NewVideoUseCase(
	videoRepo video_domain.Repository,
	storage video_domain.Storage,
	transcoder video_domain.Transcoder,
) VideoUseCaseInterface {
	return &VideoUseCase{
		videoRepo: videoRepo,
		storage:   storage,
		transcoder: transcoder,
	}
}
