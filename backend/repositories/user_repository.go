package repositories

import (
	"context"

	"github.com/TimiBolu/lema-ai-users-service/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(ctx context.Context, page, limit int) ([]models.User, int64, error)
	FindByID(ctx context.Context, id string) (*models.User, error)
	Count(ctx context.Context) (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll(ctx context.Context, page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var totalUsers int64

	offset := (page - 1) * pageSize

	if err := r.db.WithContext(ctx).Model(&models.User{}).Count(&totalUsers).Error; err != nil {
		return nil, 0, err
	}

	result := r.db.WithContext(ctx).Preload("Address").Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, totalUsers, nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).Preload("Address").First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&models.User{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
