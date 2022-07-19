package postgres_repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github/tuxoo/idler-email/internal/model/dto"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*dto.UserDTO, error) {
	var user dto.UserDTO
	query := fmt.Sprintf("SELECT id, name, email, registered_at FROM %s WHERE email=$1", userTable)
	if err := r.db.Get(&user, query, email); err != nil {
		return &user, err
	}

	return &user, nil
}
