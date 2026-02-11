package job

import (
	"time"

	"github.com/google/uuid"
)

type Metadata struct {
	ID        uuid.UUID
	Type      string
	Attempt   int
	MaxRetry  int
	CreatedAt time.Time
}
