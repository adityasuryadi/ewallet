package services

import (
	"context"

	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"github.com/adityasuryadi/ewallet/internal/models"
	"github.com/pkg/errors"
)

type LoginService struct {
	UserRepositroy interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, request *models.LoginRequest) (*models.LoginResponse, error) {
	var (
		log = helpers.Logger
	)
	user, err := s.UserRepositroy.GetUserByUsername(ctx, request.Username)
	if err != nil {
		log.Errorf("error get user %s", err.Error())
		return nil, errors.Wrap(err, "error get user")
	}

	response := &models.LoginResponse{
		Email:    user.Email,
		Username: user.Username,
	}

	return response, nil
}
