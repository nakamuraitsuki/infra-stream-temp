package outbox

import (
	"context"

	"example.com/m/internal/domain/shared"
)

func (r *outboxRepository) ListUnpublished(ctx context.Context, limit int) ([]shared.OutboxEntry, error) {
	return nil, nil
}
