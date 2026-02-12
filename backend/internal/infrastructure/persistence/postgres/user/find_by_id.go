package user

import (
	"context"

	user_domain "example.com/m/internal/domain/user"
	"github.com/google/uuid"
)

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*user_domain.User, error) {
	// TODO: implement
	return nil, nil
}
