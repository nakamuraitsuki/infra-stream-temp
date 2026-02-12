package user

import (
	"context"

	user_domain "example.com/m/internal/domain/user"
)

func (r *userRepository) Save(ctx context.Context, user *user_domain.User) error {
	// TODO: implement
	return nil
}
