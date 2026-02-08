package user

import (
	"context"

	domain "example.com/m/internal/domain/user"
	"github.com/google/uuid"
)

type UserUseCaseInterface interface {
	Register(ctx context.Context, name string, bio string) (*domain.User, error)
	UpdateProfile(ctx context.Context, id uuid.UUID, name string, bio string) error
	UpdateIcon(ctx context.Context, id uuid.UUID, data []byte) error
}

type UserUseCase struct {
	repo        domain.Repository
	iconStorage domain.IconStorage
}

func NewUserUseCase(repo domain.Repository, iconStorage domain.IconStorage) UserUseCaseInterface {
	return &UserUseCase{
		repo:        repo,
		iconStorage: iconStorage,
	}
}
