package outbox

import (
	"context"

	"example.com/m/internal/domain/event"
)

func (r *outboxRepository) Save(ctx context.Context, events []event.Event) error {
	return nil
}
