package video

import (
	"context"
	"io"
)

type Storage interface {
	SaveSource(ctx context.Context, key string, data io.Reader) error
	SaveStream(ctx context.Context, key string, data io.Reader) error
	GetStream(ctx context.Context, key string) (io.ReadSeeker, error)
	Delete(ctx context.Context, key string) error
}
