package cmd

import (
	"github.com/adityasuryadi/ewallet/bootstrap"
	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/adityasuryadi/ewallet/internal/api"
	"github.com/adityasuryadi/ewallet/internal/repository"
	"github.com/adityasuryadi/ewallet/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ServeHttp(config *viper.Viper) {
	log := helpers.Logger
	log.Info("serve http: ", config.GetString("web.port"))
	db := bootstrap.NewDatabase(config, log)

	healtCheckService := &services.HealtcheckServices{}
	userRepository := &repository.UserRepository{
		DB: db,
	}
	registerService := &services.RegisterService{
		UserRepository: userRepository,
	}
	loginService := &services.LoginService{
		UserRepositroy: userRepository,
	}
	healtcheckAPI := api.Healtcheck{HealthcheckServices: healtCheckService}
	registerAPI := api.Register{RegisterService: registerService}
	loginAPI := api.LoginHandler{
		LoginService: *loginService,
	}
	r := gin.Default()
	r.GET("/health", healtcheckAPI.HealtcheckHandlerHttp)
	r.POST("/register", registerAPI.Register)
	r.POST("/login", loginAPI.Login)

	log.Info("serve http: ", config.GetString("web.port"))
	err := r.Run(":" + config.GetString("web.port"))
	if err != nil {
		log.Fatal(err)
	}
}
