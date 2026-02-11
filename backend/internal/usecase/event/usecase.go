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

	for _, ev := range events {
		payload, err := json.Marshal(ev)
		if err != nil {
			continue // 失敗したら次は飛ばす（リトライは次回）
		}

		meta := job.Metadata{
			ID:        ev.ID(),
			Type:      ev.EventType(),
			Attempt:   0,
			MaxRetry:  3,
			CreatedAt: ev.OccurredAt(),
		}
		if err := uc.JobQueue.Enqueue(ctx, meta, payload); err != nil {
			continue // 失敗したら次は飛ばす（リトライは次回）
		}

		// 3. 送信済みマーク
		_ = uc.OutboxRepo.MarkAsPublished(ctx, ev.ID())
	}
	return nil
}
