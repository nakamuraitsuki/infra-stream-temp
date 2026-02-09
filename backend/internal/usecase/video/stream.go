package video

import (
	"context"
	"errors"
	"io"

	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

type PlaybackInfo struct {
	StreamKey string // URL 組み立てはHandler側で行う（リダイレクト）
	MIMEType  string
}

func (uc *VideoUseCase) GetPlaybackInfo(
	ctx context.Context,
	videoID uuid.UUID,
) (*PlaybackInfo, error) {
	
	video, err := uc.videoRepo.FindByID(ctx, videoID)
	if err != nil {
		return nil, err
	}

	if video.Status() != video_value.StatusReady {
		return nil, errors.New("not ready")
	}

	if video.Visibility() != video_value.VisibilityPublic {
		return nil, errors.New("forbidden")
	}

	return &PlaybackInfo{
		StreamKey: video.StreamKey(),
		MIMEType:  "application/x-mpegURL",
	}, nil
}

func (uc *VideoUseCase) GetVideoStream(
	ctx context.Context,
	videoID uuid.UUID,
) (io.ReadSeeker, string, error) {

	video, err := uc.videoRepo.FindByID(ctx, videoID)
	if err != nil {
		return nil, "", err
	}

	// Check visibility
	if video.Visibility() != video_value.VisibilityPublic {
		return nil, "", errors.New("video is not public")
	}

	// MIME type defined by application-level policy
	mimeType := "application/x-mpegURL" // HLS
	// if needed, determine MIME type based on video properties

	stream, err := uc.storage.GetStream(ctx, video.StreamKey())
	if err != nil {
		return nil, "", err
	}

	return stream, mimeType, nil
}
