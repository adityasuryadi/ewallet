package interfaces

import (
	"context"

	"github.com/adityasuryadi/ewallet/internal/models"
)

type IUserRepository interface {
	InsertUser(user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	InserUserSession(ctx context.Context, session *models.UserSession) error
	DeleteUserSession(ctx context.Context, token string) error
	GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error)
	UpdateTokenWByRefreshToken(ctx context.Context, token string, refreshToken string) error
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error)
}
