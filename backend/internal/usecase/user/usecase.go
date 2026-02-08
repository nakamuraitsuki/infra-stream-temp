package user

import (
	"context"
	"time"

	domain "example.com/m/internal/domain/user"
	"example.com/m/internal/domain/user/value"
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

func (uc *UserUseCase) Register(
	ctx context.Context,
	name string,
	bio string,
) (*domain.User, error) {

	user := domain.NewUser(
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
	data []byte,
) error {

	user, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	key := "icons/" + uuid.New().String() + ".png"

	if err := uc.iconStorage.Upload(ctx, key, data); err != nil {
		return err
	}

	user.UpdateIcon(&key)

	return uc.repo.Save(ctx, user)
}
