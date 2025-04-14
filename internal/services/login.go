package services

import (
	"context"
	"time"

	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"github.com/adityasuryadi/ewallet/internal/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepositroy interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, request *models.LoginRequest) (*models.LoginResponse, error) {
	var (
		log = helpers.Logger
		now = time.Now()
	)

	// get user by username
	user, err := s.UserRepositroy.GetUserByUsername(ctx, request.Username)
	if err != nil {
		log.Errorf("error get user %s", err.Error())
		return nil, errors.Wrap(err, "error get user")
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		log.Errorf("error compare password %s", err.Error())
		return nil, errors.Wrap(err, "failed to compare password")
	}

	// generate token
	token, err := helpers.GenerateToken(ctx, user.ID, user.Username, user.Fullname, user.Email, time.Now(), "token")
	if err != nil {
		log.Errorf("error generate token %s", err.Error())
		return nil, errors.Wrap(err, "failed to generate token")
	}

	// generate refresh token
	refreshToken, err := helpers.GenerateToken(ctx, user.ID, user.Username, user.Fullname, user.Email, time.Now(), "refresh_token")
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate refresh token")
	}

	// insert user session
	userSession := &models.UserSession{
		UserID:              user.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}

	err = s.UserRepositroy.InserUserSession(ctx, userSession)
	if err != nil {
		log.Errorf("error insert user session %s", err.Error())
		return nil, errors.Wrap(err, "failed to insert user session")
	}

	response := &models.LoginResponse{
		Email:        user.Email,
		Username:     user.Username,
		Fullname:     user.Fullname,
		Token:        token,
		RefreshToken: refreshToken,
	}

	return response, nil
}
