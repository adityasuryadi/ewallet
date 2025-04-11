package repository

import (
	"context"

	"github.com/adityasuryadi/ewallet/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertUser(user *models.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user := new(models.User)
	err := r.DB.Where("username = ?", "aditya").First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
