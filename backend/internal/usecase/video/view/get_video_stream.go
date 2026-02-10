package view

import (
	"context"
	"io"

	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

func (uc *VideoViewingUseCase) GetVideoStream(
	ctx context.Context,
	videoID uuid.UUID,
) (io.ReadSeeker, string, error) {

	video, err := uc.VideoRepo.FindByID(ctx, videoID)
	if err != nil {
		return nil, "", err
	}

	if video.Status() != video_value.StatusReady {
		return nil, "", ErrVideoNotReady
	}

	// Check visibility
	if video.Visibility() != video_value.VisibilityPublic {
		return nil, "", ErrVideoForbidden
	}

	// MIME type defined by application-level policy
	mimeType := "application/x-mpegURL" // HLS
	// if needed, determine MIME type based on video properties

	stream, err := uc.Storage.GetStream(ctx, video.StreamKey())
	if err != nil {
		return nil, "", err
	}

	return stream, mimeType, nil
}
