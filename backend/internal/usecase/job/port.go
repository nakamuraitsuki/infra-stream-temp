package job

import "context"

type Queue interface {
	Enqueue(ctx context.Context, meta Metadata, payload []byte) error
}

type Consumer interface {
	Start(ctx context.Context, handler Handler) error
}