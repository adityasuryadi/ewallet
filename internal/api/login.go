package api

import (
	"errors"
	"net/http"

	"github.com/adityasuryadi/ewallet/constants"
	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/adityasuryadi/ewallet/internal/models"
	"github.com/adityasuryadi/ewallet/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginHandler struct {
	LoginService services.LoginService
}

func (api *LoginHandler) Login(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	request := new(models.LoginRequest)
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("error bind json ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if err := request.Validate(); err != nil {
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	response, err := api.LoginService.Login(c.Request.Context(), request)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		helpers.SendResponseHTTP(c, 404, "user not found", nil)
		return
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("error login ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, response)
}
