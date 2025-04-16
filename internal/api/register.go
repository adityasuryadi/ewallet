package api

import (
	"net/http"

	"github.com/adityasuryadi/ewallet/constants"
	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"github.com/adityasuryadi/ewallet/internal/models"
	"github.com/gin-gonic/gin"
)

type Register struct {
	RegisterService interfaces.IRegisterService
}

func (api *Register) Register(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	req := models.User{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("error bind json ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	resp, err := api.RegisterService.Register(c, req)
	if err != nil {
		log.Error("error register ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
}
