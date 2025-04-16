package repository

import (
	"context"
	"errors"
	"fmt"

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

func (r *UserRepository) DeleteUserSession(ctx context.Context, token string) error {
	err := r.DB.Where("token = ?", token).Delete(&models.UserSession{}).Debug().Error
	if err != nil {
		fmt.Printf("error delete user session %s", err.Error())
		return err
	}
	return nil
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error) {
	var (
		session models.UserSession
		err     error
	)
	err = r.DB.Where("token = ?", token).First(&session).Error
	if err != nil {
		return session, err
	}
	if session.ID == 0 {
		return session, errors.New("session not found")
	}
	return session, nil
}

func (r *UserRepository) GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error) {
	var (
		session models.UserSession
		err     error
	)
	err = r.DB.Where("refresh_token = ?", refreshToken).First(&session).Error
	if err != nil {
		return session, err
	}
	if session.ID == 0 {
		return session, errors.New("session not found")
	}
	return session, nil
}

func (r *UserRepository) UpdateTokenWByRefreshToken(ctx context.Context, token string, refreshToken string) error {
	return r.DB.Exec("UPDATE user_sessions SET token = ? WHERE refresh_token = ?", token, refreshToken).Error
}
