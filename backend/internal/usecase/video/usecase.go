package video

import "errors"

type VideoUseCaseInterface interface {
	VideoManagementUseCaseInterface
	VideoViewingUseCaseInterface
}

var (
	ErrVideoNotReady  = errors.New("video is not ready for playback")
	ErrVideoForbidden = errors.New("video is not accessible")
)
