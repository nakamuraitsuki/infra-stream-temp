package event

import (
	"context"

	"example.com/m/internal/domain/shared"
	"example.com/m/internal/usecase/job"
	"example.com/m/internal/usecase/tx"
)

type EventRelayUseCaseInterface interface {
	ProcessOutbox(ctx context.Context) error
}

type EventRelayUseCase struct {
	OutboxRepo shared.OutboxRepository
	JobQueue   job.Queue
	UoW        tx.UnitOfWork
}

func NewEventRelayUseCase(
	outboxRepo shared.OutboxRepository,
	jobQueue job.Queue,
	uow tx.UnitOfWork,
) EventRelayUseCaseInterface {
	return &EventRelayUseCase{
		OutboxRepo: outboxRepo,
		JobQueue:   jobQueue,
		UoW:        uow,
	}
}

func (uc *EventRelayUseCase) ProcessOutbox(ctx context.Context) error {
	return uc.UoW.Do(ctx, func(ctx context.Context) error {

		entries, err := uc.OutboxRepo.ListUnpublished(ctx, 10)
		if err != nil {
			return err
		}

		var firstErr error

		for _, entry := range entries {

			meta := job.Metadata{
				ID:        entry.ID,
				Type:      entry.EventType,
				Attempt:   0,
				MaxRetry:  3,
				CreatedAt: entry.OccurredAt,
			}

			if err := uc.JobQueue.Enqueue(ctx, meta, entry.Payload); err != nil {
				if firstErr == nil {
					firstErr = err
				}
				continue
			}

			if err := uc.OutboxRepo.MarkAsPublished(ctx, entry.ID); err != nil {
				// これは致命的
				return err
			}
		}

		return firstErr
	})
}
