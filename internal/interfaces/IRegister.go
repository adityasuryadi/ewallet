package interfaces

import (
	"context"

	"github.com/adityasuryadi/ewallet/internal/models"
	"github.com/gin-gonic/gin"
)

type IRegisterHandler interface {
	Register(c *gin.Context)
}

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
