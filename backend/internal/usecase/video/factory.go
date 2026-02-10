package video

import (
	video_domain "example.com/m/internal/domain/video"
	"example.com/m/internal/usecase/video/manage"
	"example.com/m/internal/usecase/video/view"
)

func NewVideoUseCase(
	videoRepo video_domain.Repository,
	storage video_domain.Storage,
	transcoder video_domain.Transcoder,
) VideoUseCaseInterface {
	return &struct {
		manage.VideoManagementUseCaseInterface
		view.VideoViewingUseCaseInterface
	}{
		VideoManagementUseCaseInterface: &manage.VideoManagementUseCase{
			VideoRepo:  videoRepo,
			Storage:    storage,
			Transcoder: transcoder,
		},
		VideoViewingUseCaseInterface: &view.VideoViewingUseCase{
			VideoRepo: videoRepo,
			Storage:   storage,
		},
	}
}
