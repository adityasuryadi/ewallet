package interfaces

import "github.com/gin-gonic/gin"

type IHealthcheckService interface {
	HealthcheckServices() (string, error)
}

type IHealthcheckHandler interface {
	Healtcheck(c *gin.Context)
}
