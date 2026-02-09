package user

import "context"

type IconStorage interface {
	Upload(ctx context.Context, key string, data []byte) error
}
