package video

import (
	"context"
	"io"

	"github.com/google/uuid"
)

func (uc *VideoUseCase) GetVideoStream(
	ctx context.Context,
	videoID uuid.UUID,
) (io.ReadSeeker, string, error) {
	// TODO: 実装
	return nil, "", nil
}
