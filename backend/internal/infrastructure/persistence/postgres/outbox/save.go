package outbox

import (
	"context"
	"fmt"

	"example.com/m/internal/domain/event"
	"example.com/m/internal/infrastructure/persistence/postgres"
	"github.com/jmoiron/sqlx"
)

func (r *outboxRepository) Save(ctx context.Context, events []event.Event) error {
	if len(events) == 0 {
		return nil
	}
	db := postgres.GetExt(ctx, r.db)

	dtos := make([]*outboxDTO, len(events))
	for i, e := range events {
		m, err := fromEntity(e)
		if err != nil {
			return fmt.Errorf("failed to map event to outbox DTO: %w", err)
		}
		dtos[i] = m
	}

	const query = `
INSERT INTO outbox (id, event_type, payload, occurred_at)
VALUES (?)	
`
	q, args, err := sqlx.In(query, dtos)
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	// NOTE: sqlx.In はプレースホルダを "?" で生成する
	// 			 PostgreSQL では "$1", "$2", ... の形式が必要なため、Rebind 
	q = r.db.Rebind(q)

	if _, err := db.ExecContext(ctx, q, args...); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
