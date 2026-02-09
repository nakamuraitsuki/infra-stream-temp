package user

import (
	"context"

	"example.com/m/internal/domain/user"
	"github.com/google/uuid"
)

func (u *UserUseCase) GetMe(ctx context.Context, id uuid.UUID) (*user.User, error) {
	return u.repo.FindByID(ctx, id)
}
