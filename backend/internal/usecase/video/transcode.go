package video

import (
	"context"

	"github.com/google/uuid"
)

func (uc *VideoUseCase) StartTranscoding(
	ctx context.Context,
	videoID uuid.UUID,
	streamKey string,
) error {
	// TODO: 実装
	return nil
}
