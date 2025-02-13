package repositories

import (
	"context"

	"github.com/TimiBolu/lema-ai-users-service/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindByUserID(ctx context.Context, userID string) ([]models.Post, error)
	Create(ctx context.Context, post *models.Post) error
	Delete(ctx context.Context, id string) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) FindByUserID(ctx context.Context, userID string) ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) Create(ctx context.Context, post *models.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *postRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Post{}).Error
}
