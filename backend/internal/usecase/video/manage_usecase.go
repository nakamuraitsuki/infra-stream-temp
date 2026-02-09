package video

import (
	"context"
	"io"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

type VideoManagementUseCaseInterface interface {
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
}

type videoManagementUseCase struct {
	videoRepo  video_domain.Repository
	storage    video_domain.Storage
	transcoder video_domain.Transcoder
}
