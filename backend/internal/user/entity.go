package user

import (
	"time"

	"example.com/m/internal/user/value"
	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	name      string
	bio       string
	iconKey   *string
	role      value.Role
	createdAt time.Time
}

func NewUser(
	id uuid.UUID,
	name string,
	bio string,
	iconKey *string,
	role value.Role,
	createdAt time.Time,
) *User {
	return &User{
		id:        id,
		name:      name,
		bio:       bio,
		iconKey:   iconKey,
		role:      role,
		createdAt: createdAt,
	}
}
