package video

import (
	"context"
	"io"

	"github.com/google/uuid"
)

func (uc *VideoUseCase) UploadSource(
	ctx context.Context,
	videoID uuid.UUID,
	videoData io.Reader,
) error {
	// TODO: 実装
	return nil
}