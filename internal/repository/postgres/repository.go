package postgres_repository

import (
	"github.com/jmoiron/sqlx"
	"github/eugene-krivtsov/idler-email/internal/model/dto"
)

const (
	userTable = "\"user\""
)

type Users interface {
	FindByEmail(email string) (*dto.UserDTO, error)
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: NewUserRepository(db),
	}
}
