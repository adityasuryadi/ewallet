package repository

import (
	"github.com/adityasuryadi/ewallet/internal/models"
	"gorm.io/gorm"
)

type RegisterRepository struct {
	DB *gorm.DB
}

func (r *RegisterRepository) InsertUser(user *models.User) error {
	return r.DB.Create(&user).Error
}
