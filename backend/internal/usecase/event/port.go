package event

import "context"

type Relay interface {
	Start(ctx context.Context) error
}