package outbox

import (
	"encoding/json"
	"time"

	"example.com/m/internal/domain/event"
	"example.com/m/internal/domain/shared"
	"github.com/google/uuid"
)

type outboxDTO struct {
	ID         uuid.UUID `db:"id"`
	EventType  string    `db:"event_type"`
	Payload    []byte    `db:"payload"`
	OccurredAt time.Time `db:"occurred_at"`
}

func fromEntity(e event.Event) (*outboxDTO, error) {
	payload, err := json.Marshal(e.Payload())
	if err != nil {
		return nil, err
	}

	return &outboxDTO{
		ID:         e.ID(),
		EventType:  e.EventType(),
		Payload:    payload,
		OccurredAt: e.OccurredAt(),
	}, nil
}

func (dto *outboxDTO) toEntry() (shared.OutboxEntry, error) {
	return shared.OutboxEntry{
		ID:         dto.ID,
		EventType:  dto.EventType,
		Payload:    dto.Payload,
		OccurredAt: dto.OccurredAt,
	}, nil
}