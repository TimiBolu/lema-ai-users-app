package services

import (
	"context"

	"github.com/TimiBolu/lema-ai-users-service/models"
	"github.com/TimiBolu/lema-ai-users-service/repositories"
)

type UserService interface {
	GetUsers(ctx context.Context, page, size int) ([]models.User, int64, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserCount(ctx context.Context) (int64, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUsers(ctx context.Context, page, size int) ([]models.User, int64, error) {
	return s.repo.FindAll(ctx, page, size)
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *userService) GetUserCount(ctx context.Context) (int64, error) {
	return s.repo.Count(ctx)
}
