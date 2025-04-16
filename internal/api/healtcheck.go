package api

import (
	"net/http"

	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"github.com/gin-gonic/gin"
)

type HealtcheckHandler struct {
	HealthcheckServices interfaces.IHealthcheckService
}

func (api *HealtcheckHandler) Healtcheck(c *gin.Context) {
	msg, err := api.HealthcheckServices.HealthcheckServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, msg)
	return
}

func (api *HealtcheckHandler) HealtcheckHandlerHttp(c *gin.Context) {
	msg, err := api.HealthcheckServices.HealthcheckServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, msg)
}
