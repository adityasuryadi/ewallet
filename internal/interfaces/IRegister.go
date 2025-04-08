package interfaces

import (
	"context"

	"github.com/adityasuryadi/ewallet/internal/models"
)

type IRegisterRepository interface {
	InsertUser(user *models.User) error
}

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
