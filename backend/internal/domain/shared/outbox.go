package shared

import (
	"context"

	"example.com/m/internal/domain/event"
	"github.com/google/uuid"
)

type OutboxRepository interface {
	Save(ctx context.Context, event event.Event) error

	ListUnpublished(ctx context.Context, limit int) ([]event.Event, error)

	MarkAsPublished(ctx context.Context, id uuid.UUID) error
}
