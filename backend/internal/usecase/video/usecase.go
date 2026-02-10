package video

import (
	"example.com/m/internal/usecase/video/manage"
	"example.com/m/internal/usecase/video/view"
)

type VideoUseCaseInterface interface {
	manage.VideoManagementUseCaseInterface
	view.VideoViewingUseCaseInterface
}
