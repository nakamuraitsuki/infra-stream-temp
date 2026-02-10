package view

import (
	"context"
	"io"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"example.com/m/internal/usecase/video/query"
	"github.com/google/uuid"
)

type PlaybackInfo struct {
	StreamKey string // URL 組み立てはHandler側で行う（リダイレクト）
	MIMEType  string
}

func (uc *VideoViewingUseCase) GetByID(
	ctx context.Context,
	videoID uuid.UUID,
) (*video_domain.Video, error) {

	video, err := uc.VideoRepo.FindByID(ctx, videoID)
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (uc *VideoViewingUseCase) ListPublic(
	ctx context.Context,
	query query.VideoSearchQuery,
) ([]*video_domain.Video, error) {

	visibility := video_value.VisibilityPublic
	status := video_value.StatusReady

	cond := video_domain.ListCondition{
		Visibility: &visibility,
		Status:     &status,
		Limit:      query.Limit,
	}

	return uc.VideoRepo.FindByCondition(ctx, cond)
}

func (uc *VideoViewingUseCase) SearchByTag(
	ctx context.Context,
	tag video_value.Tag,
	query query.VideoSearchQuery,
) ([]*video_domain.Video, error) {

	visibility := video_value.VisibilityPublic
	status := video_value.StatusReady

	cond := video_domain.ListCondition{
		Tag:        &tag,
		Visibility: &visibility,
		Status:     &status,
		Limit:      query.Limit,
	}

	return uc.VideoRepo.FindByCondition(ctx, cond)
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
