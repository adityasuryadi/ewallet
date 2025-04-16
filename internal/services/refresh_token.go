package services

import (
	"context"
	"time"

	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"github.com/adityasuryadi/ewallet/internal/models"
	"github.com/pkg/errors"
)

type RefreshTokenService struct {
	UserRepository interfaces.IUserRepository
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error) {
	resp := models.RefreshTokenResponse{}
	token, err := helpers.GenerateToken(ctx, tokenClaim.UserId, tokenClaim.Username, tokenClaim.Fullname, tokenClaim.Email, time.Now(), "token")
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate new token")
	}

	err = s.UserRepository.UpdateTokenWByRefreshToken(ctx, token, refreshToken)
	if err != nil {
		return resp, errors.Wrap(err, "failed to update new token")
	}
	resp.Token = token
	return resp, nil
}
