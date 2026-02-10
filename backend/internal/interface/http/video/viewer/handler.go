package viewer

import "example.com/m/internal/usecase/video/view"

type VideoViewingHandler struct {
	usecase view.VideoViewingUseCaseInterface
}

func NewVideoViewingHandler(usecase *view.VideoViewingUseCase) *VideoViewingHandler {
	return &VideoViewingHandler{
		usecase: usecase,
	}
}
