package services

import (
	"context"
	"errors"

	"github.com/TimiBolu/lema-ai-users-service/models"
	"github.com/TimiBolu/lema-ai-users-service/repositories"
	"github.com/google/uuid"
)

type PostService interface {
	GetPostsByUser(ctx context.Context, userID string) ([]models.Post, error)
	CreatePost(ctx context.Context, title, body, userID string) (*models.Post, error)
	DeletePost(ctx context.Context, id string) error
}

type postService struct {
	repo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postService{repo}
}

func (s *postService) GetPostsByUser(ctx context.Context, userID string) ([]models.Post, error) {
	if userID == "" {
		return nil, errors.New("userID is required")
	}
	return s.repo.FindByUserID(ctx, userID)
}

func (s *postService) CreatePost(ctx context.Context, title, body, userID string) (*models.Post, error) {
	if title == "" || body == "" || userID == "" {
		return nil, errors.New("missing required fields")
	}
	newPost := &models.Post{
		ID:     uuid.NewString(),
		Title:  title,
		Body:   body,
		UserID: userID,
	}
	if err := s.repo.Create(ctx, newPost); err != nil {
		return nil, err
	}
	return newPost, nil
}

func (s *postService) DeletePost(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
