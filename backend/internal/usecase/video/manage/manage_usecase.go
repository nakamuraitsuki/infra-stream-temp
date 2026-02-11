package manage

import (
	"context"
	"io"

	"example.com/m/internal/domain/shared"
	video_domain "example.com/m/internal/domain/video"
	"example.com/m/internal/usecase/tx"
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
		tags []string,
	) (*CreateResponse, error)

	// UploadSource uploads the raw video data for the specified video.
	UploadSource(ctx context.Context, videoID uuid.UUID, videoData io.Reader) error

	// StartTranscoding initiates the transcoding process for a video.
	StartTranscoding(ctx context.Context, videoID uuid.UUID) error

	// ListMine returns a list of videos owned by the specified user.
	ListMine(ctx context.Context, ownerID uuid.UUID, query query.VideoSearchQuery) (*ListMineResults, error)
}

type VideoManagementUseCase struct {
	VideoRepo  video_domain.Repository
	OutboxRepo shared.OutboxRepository
	Storage    video_domain.Storage
	Transcoder video_domain.Transcoder
	UoW        tx.UnitOfWork
}

func NewVideoManagementUseCase(
	videoRepo video_domain.Repository,
	outboxRepo shared.OutboxRepository,
	storage video_domain.Storage,
	transcoder video_domain.Transcoder,
	uow tx.UnitOfWork,
) VideoManagementUseCaseInterface {
	return &VideoManagementUseCase{
		VideoRepo:  videoRepo,
		OutboxRepo: outboxRepo,
		Storage:    storage,
		Transcoder: transcoder,
		UoW:        uow,
	}
}
