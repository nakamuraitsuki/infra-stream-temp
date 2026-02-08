package user

import (
	"context"
	"time"

	domain "example.com/m/internal/domain/user"
	"example.com/m/internal/domain/user/value"
	"github.com/google/uuid"
)

type UseCase interface {
	Register(ctx context.Context, name string, bio string) (*domain.User, error)
	UpdateProfile(ctx context.Context, id uuid.UUID, name string, bio string) error
	UpdateIcon(ctx context.Context, id uuid.UUID, iconKey *string) error
}

type UserUseCase struct {
	repo domain.Repository
}

func NewUserUseCase(repo domain.Repository) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (uc *UserUseCase) Register(
	ctx context.Context,
	name string,
	bio string,
) (*domain.User, error) {

	user :=  domain.NewUser(
		uuid.New(),
		name,
		bio,
		nil,
		value.RoleUser,
		time.Now(),
	)

	if err := uc.repo.Save(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) UpdateProfile(
	ctx context.Context,
	id uuid.UUID,
	name string,
	bio string,
) error {

	user, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err := user.UpdateProfile(name, bio); err != nil {
		return err
	}

	return uc.repo.Save(ctx, user)
}

func (uc *UserUseCase) UpdateIcon(
	ctx context.Context,
	id uuid.UUID,
	iconKey *string,
) error {

	user, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	user.UpdateIcon(iconKey)

	return uc.repo.Save(ctx, user)
}
