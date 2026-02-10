package view

import (
	"context"

	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

type PlaybackInfo struct {
	StreamKey string // URL 組み立てはHandler側で行う（リダイレクト）
	MIMEType  string
}

func (uc *VideoViewingUseCase) GetPlaybackInfo(
	ctx context.Context,
	videoID uuid.UUID,
) (*PlaybackInfo, error) {

	video, err := uc.VideoRepo.FindByID(ctx, videoID)
	if err != nil {
		return nil, err
	}

	if video.Status() != video_value.StatusReady {
		return nil, ErrVideoNotReady
	}

	if video.Visibility() != video_value.VisibilityPublic {
		return nil, ErrVideoForbidden
	}

	return &PlaybackInfo{
		StreamKey: video.StreamKey(),
		MIMEType:  "application/x-mpegURL",
	}, nil
}
