package api

import (
	"net/http"

	"github.com/adityasuryadi/ewallet/constants"
	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (api *LogoutHandler) Logout(c *gin.Context) {
	err := api.LogoutService.Logout(c, c.GetString("token"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, nil)
}
