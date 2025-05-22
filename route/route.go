package route

import (
	"sistem-pembiayaan/controller"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Router                *gin.Engine
	UserController        *controller.UserController
	CalculationController *controller.CalculationController
}

func (rc *RouteConfig) Setup() {
	api := rc.Router.Group("/api/v1")

	// User routes
	userRoute := &UserRoute{
		UserController: rc.UserController,
	}
	userRoute.RegisterRoutes(api)

	// Calculation routes
	calcRoute := &CalculationRoute{
		CalculationController: rc.CalculationController,
	}
	calcRoute.CalculationRoutes(api)
}
