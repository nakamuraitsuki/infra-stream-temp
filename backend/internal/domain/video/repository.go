package video

import (
	"context"

	"example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Video, error)
	Save(ctx context.Context, v *Video) error

	UpdateProcessingResult(
		ctx context.Context,
		id uuid.UUID,
		streamKey string,
		status value.Status,
		retryCount int,
		failureReason *value.FailureReason,
	) error

	ListByOwner(
		ctx context.Context,
		ownerID uuid.UUID,
		limit int,
	) ([]*Video, error)

	ListPublic(
		ctx context.Context,
		limit int,
	) ([]*Video, error)

	SearchByTag(
		ctx context.Context,
		tag value.Tag,
		limit int,
	) ([]*Video, error)
}
