package video

import (
	"context"
	"io"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

type VideoUseCaseInterface interface {
	// Create uploads a new video.
	Create(
		ctx context.Context,
		ownerID uuid.UUID,
		sourceKey string,
		title string,
		description string,
		tags []video_value.Tag,
		visibility video_value.Visibility,
	) (*video_domain.Video, error)

	// ListMine returns a list of videos owned by the specified user.
	ListMine(ctx context.Context, ownerID uuid.UUID, query VideoSearchQuery) ([]*video_domain.Video, error)

	// ListPublic returns a list of publicly available videos for viewing.
	ListPublic(ctx context.Context, query VideoSearchQuery) ([]*video_domain.Video, error)

	// SearchByTag returns a list of videos matching the specified tag.
	SearchByTag(ctx context.Context, tag video_value.Tag, query VideoSearchQuery) ([]*video_domain.Video, error)

	// StartTranscoding initiates the transcoding process for a video.
	StartTranscoding(ctx context.Context, videoID uuid.UUID, streamKey string) error

	// GetByID returns a video by its ID.
	GetByID(ctx context.Context, videoID uuid.UUID) (*video_domain.Video, error)

	// GetVideoStream returns a readable video stream and its MIME type.
	//
	// The returned io.ReadSeeker represents the video content to be streamed.
	// The caller (e.g. HTTP handler) is responsible for closing it if necessary.
	//
	// The MIME type is determined by application-level policy
	// (e.g. video/mp4, application/x-mpegURL),
	// not by HTTP-specific concerns.
	GetVideoStream(ctx context.Context, videoID uuid.UUID) (io.ReadSeeker, string, error)
}

type VideoUseCase struct {
	videoRepo video_domain.Repository
}

func NewVideoUseCase(videoRepo video_domain.Repository) VideoUseCaseInterface {
	return &VideoUseCase{
		videoRepo: videoRepo,
	}
}
