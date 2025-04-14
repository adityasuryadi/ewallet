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
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) InserUserSession(ctx context.Context, session *models.UserSession) error {
	err := r.DB.Create(&session).Error
	if err != nil {
		return err
	}
	return nil
}
