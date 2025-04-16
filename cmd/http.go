package cmd

import (
	"log"

	"github.com/adityasuryadi/ewallet/bootstrap"
	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/adityasuryadi/ewallet/internal/api"
	"github.com/adityasuryadi/ewallet/internal/interfaces"
	"github.com/adityasuryadi/ewallet/internal/repository"
	"github.com/adityasuryadi/ewallet/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Dependency struct {
	UserRepository interfaces.IUserRepository
	RegisterAPI    interfaces.IRegisterHandler
	LoginAPI       interfaces.ILoginHandler
	LogoutAPI      interfaces.ILogoutHandler
	HealthcheckAPI interfaces.IHealthcheckHandler
}

func dependencyInject(config *viper.Viper) Dependency {
	log := helpers.Logger
	log.Info("serve http: ", config.GetString("web.port"))
	db := bootstrap.NewDatabase(config, log)

	userRepository := &repository.UserRepository{
		DB: db,
	}
	registerService := &services.RegisterService{
		UserRepository: userRepository,
	}
	loginService := &services.LoginService{
		UserRepositroy: userRepository,
	}
	logoutService := &services.LogoutService{
		UserRepository: userRepository,
	}
	healthcheckAPI := api.HealtcheckHandler{}
	registerAPI := api.Register{RegisterService: registerService}
	loginAPI := api.LoginHandler{
		LoginService: *loginService,
	}

	logoutAPI := api.LogoutHandler{
		LogoutService: logoutService,
	}
	return Dependency{
		UserRepository: userRepository,
		RegisterAPI:    &registerAPI,
		LoginAPI:       &loginAPI,
		LogoutAPI:      &logoutAPI,
		HealthcheckAPI: &healthcheckAPI,
	}
}

func ServeHttp(config *viper.Viper) {

	dependency := dependencyInject(config)
	r := gin.Default()
	r.GET("/health", dependency.HealthcheckAPI.Healtcheck)
	r.POST("/register", dependency.RegisterAPI.Register)
	r.POST("/login", dependency.LoginAPI.Login)
	r.POST("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutAPI.Logout)
	err := r.Run(":" + config.GetString("web.port"))
	if err != nil {
		log.Fatal(err)
	}
}
