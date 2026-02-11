package event

import (
	"context"
	"encoding/json"

	"example.com/m/internal/domain/shared"
	"example.com/m/internal/usecase/job"
)

type EventRelayUseCaseInterface interface {
	ProcessOutbox(ctx context.Context) error
}

type EventRelayUseCase struct {
	OutboxRepo shared.OutboxRepository
	JobQueue   job.Queue
}

func NewEventRelayUseCase(
	outboxRepo shared.OutboxRepository,
	jobQueue job.Queue,
) EventRelayUseCaseInterface {
	return &EventRelayUseCase{
		OutboxRepo: outboxRepo,
		JobQueue:   jobQueue,
	}
}

func (uc *EventRelayUseCase) ProcessOutbox(ctx context.Context) error {
	events, err := uc.OutboxRepo.ListUnpublished(ctx, 10)
	if err != nil {
		return err
	}

	var firstErr error

	for _, ev := range events {

		payload, err := json.Marshal(ev.Payload())
		if err != nil {
			if firstErr == nil {
				firstErr = err
			}
			continue
		}

		meta := job.Metadata{
			ID:        ev.ID(),
			Type:      ev.EventType(),
			Attempt:   0,
			MaxRetry:  3,
			CreatedAt: ev.OccurredAt(),
		}

		if err := uc.JobQueue.Enqueue(ctx, meta, payload); err != nil {
			if firstErr == nil {
				firstErr = err
			}
			continue
		}

		if err := uc.OutboxRepo.MarkAsPublished(ctx, ev.ID()); err != nil {
			// これは致命的
			return err
		}
	}

	return firstErr
}

