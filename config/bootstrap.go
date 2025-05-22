package config

import (
	"sistem-pembiayaan/controller"
	"sistem-pembiayaan/repository"
	"sistem-pembiayaan/route"
	"sistem-pembiayaan/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	Router   *gin.Engine
	Validate *validator.Validate
	Viper    *viper.Viper
}

func Bootstrap(c *BootstrapConfig) {
	// Repository
	userRepo := repository.NewUserRepository(c.DB)
	tenorRepo := repository.NewTenorRepository(c.DB)

	// Service
	userService := service.NewUserService(userRepo)
	calcService := service.NewCalculationService(tenorRepo)

	// Controller
	userController := controller.NewUserController(userService, c.Validate)
	calcController := controller.NewCalculationController(calcService)

	// Routes
	routeConfig := route.RouteConfig{
		Router:                c.Router,
		UserController:        userController,
		CalculationController: calcController,
	}
	routeConfig.Setup()
}
