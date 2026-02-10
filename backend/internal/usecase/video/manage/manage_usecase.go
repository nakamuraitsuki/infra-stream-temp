package manage

import (
	"context"
	"io"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"example.com/m/internal/usecase/video/query"
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
	ListMine(ctx context.Context, ownerID uuid.UUID, query query.VideoSearchQuery) ([]*video_domain.Video, error)
}

type VideoManagementUseCase struct {
	VideoRepo  video_domain.Repository
	Storage    video_domain.Storage
	Transcoder video_domain.Transcoder
}
