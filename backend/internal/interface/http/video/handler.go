package video

import video_usecase "example.com/m/internal/usecase/video"

type VideoHandler struct {
	videoUsecase video_usecase.VideoUseCaseInterface
}

func NewVideoHandler(usecase video_usecase.VideoUseCaseInterface) *VideoHandler {
	return &VideoHandler{
		videoUsecase: usecase,
	}
}