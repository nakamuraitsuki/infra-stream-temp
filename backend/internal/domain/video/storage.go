package video

import "context"

type Storage interface {
	SaveSource(ctx context.Context, key string, data []byte) error
	SaveStream(ctx context.Context, key string, data []byte) error
	Delete(ctx context.Context, key string) error
}
