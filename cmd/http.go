package cmd

import (
	"log"

	"github.com/adityasuryadi/ewallet/internal/api"
	"github.com/adityasuryadi/ewallet/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ServeHttp(config *viper.Viper) {
	healtCheckService := &services.HealtcheckServices{}
	registerService := &services.RegisterService{}
	healtcheckAPI := api.Healtcheck{HealthcheckServices: healtCheckService}
	registerAPI := api.Register{RegisterService: registerService}
	r := gin.Default()
	r.GET("/health", healtcheckAPI.HealtcheckHandlerHttp)
	r.POST("/register", registerAPI.Register)

	err := r.Run(":" + config.GetString("web.port"))
	if err != nil {
		log.Fatal(err)
	}
}
