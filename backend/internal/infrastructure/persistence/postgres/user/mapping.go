package user

import (
	"time"

	user_domain "example.com/m/internal/domain/user"
	user_value "example.com/m/internal/domain/user/value"
	"github.com/google/uuid"
)

type userDAO struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Bio       string    `db:"bio"`
	IconKey   *string   `db:"icon_key"`
	Role      string    `db:"role"` // value.Role を文字列として保存
	CreatedAt time.Time `db:"created_at"`
}

func fromEntity(u *user_domain.User) *userDAO {
	return &userDAO{
		ID:        u.ID(),
		Name:      u.Name(),
		Bio:       u.Bio(),
		IconKey:   u.IconKey(),
		Role:      string(u.Role()),
		CreatedAt: u.CreatedAt(),
	}
}

// toEntity は Userのコンストラクタのラッパー
// DB特有の不正値のチェックなどをしたい場合に追記する
func (dao *userDAO) toEntity() (*user_domain.User, error) {
	return user_domain.NewUser(
		dao.ID,
		dao.Name,
		dao.Bio,
		dao.IconKey,
		user_value.Role(dao.Role),
		dao.CreatedAt,
	), nil
}
