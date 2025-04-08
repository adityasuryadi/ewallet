package api

import (
	"net/http"

	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"github.com/gin-gonic/gin"
)

type Healtcheck struct {
	HealthcheckServices interfaces.IHealthcheckService
}

func (api *Healtcheck) Healtcheck(c *gin.Context) error {
	msg, err := api.HealthcheckServices.HealthcheckServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, msg)
	return nil
}

func (api *Healtcheck) HealtcheckHandlerHttp(c *gin.Context) {
	msg, err := api.HealthcheckServices.HealthcheckServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, msg)
}
