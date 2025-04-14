package interfaces

import (
	"context"

	"github.com/adityasuryadi/ewallet/internal/models"
)

type IUserRepository interface {
	InsertUser(user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	InserUserSession(ctx context.Context, session *models.UserSession) error
}

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
