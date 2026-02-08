package user

import (
	"context"
	"time"

	domain "example.com/m/internal/domain/user"
	"example.com/m/internal/domain/user/value"
	"github.com/google/uuid"
)

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
