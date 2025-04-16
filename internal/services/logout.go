package services

import (
	"context"
	"errors"

	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"gorm.io/gorm"
)

type LogoutService struct {
	UserRepository interfaces.IUserRepository
}

func (s *LogoutService) Logout(ctx context.Context, token string) error {
	var (
		log = helpers.Logger
	)
	err := s.UserRepository.DeleteUserSession(ctx, token)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorf("error delete user session %s", err.Error())
		return err
	}

	return nil
}
