package user

import (
	"github.com/jmoiron/sqlx"
	user_domain "example.com/m/internal/domain/user"
)

// 小文字で開始し、このパッケージ内に閉じ込める
type userRepository struct {
	db *sqlx.DB
}

// 戻り値はドメイン層のインターフェース
func NewRepository(db *sqlx.DB) user_domain.Repository {
	return &userRepository{
		db: db,
	}
}