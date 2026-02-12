package outbox

import (
	"context"

	"github.com/google/uuid"
)

func (r *outboxRepository) MarkAsPublished(ctx context.Context, id uuid.UUID) error {
	return nil
}
