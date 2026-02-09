package video

import video_domain "example.com/m/internal/domain/video"

func NewVideoUseCase(
	videoRepo video_domain.Repository,
	storage video_domain.Storage,
	transcoder video_domain.Transcoder,
) VideoUseCaseInterface {
	return &struct {
		VideoManagementUseCaseInterface
		VideoViewingUseCaseInterface
	} {
		VideoManagementUseCaseInterface: &videoManagementUseCase{
			videoRepo:  videoRepo,
			storage:    storage,
			transcoder: transcoder,
		},
		VideoViewingUseCaseInterface: &videoViewingUseCase{
			videoRepo: videoRepo,
			storage:   storage,
		},
	}
}
