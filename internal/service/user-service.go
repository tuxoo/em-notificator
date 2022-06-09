package service

import (
	"context"
	"github/eugene-krivtsov/idler-email/internal/model/dto"
	postgres_repository "github/eugene-krivtsov/idler-email/internal/repository/postgres"
)

type UserService struct {
	repository postgres_repository.Users
}

func NewUserService(repository postgres_repository.Users) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*dto.UserDTO, error) {
	return s.repository.FindByEmail(email)
}
