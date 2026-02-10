package manager

import (
	"example.com/m/internal/usecase/video/manage"
	"example.com/m/internal/usecase/video/view"
)

type VideoManagementHandler struct {
	manageUsecase manage.VideoManagementUseCaseInterface
}

func NewVideoManagementHandler(manageUsecase *manage.VideoManagementUseCase, viewUsecase *view.VideoViewingUseCase) *VideoManagementHandler {
	return &VideoManagementHandler{
		manageUsecase: manageUsecase,
	}
}
