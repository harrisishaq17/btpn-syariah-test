package route

import (
	"sistem-pembiayaan/controller"

	"github.com/gin-gonic/gin"
)

type CalculationRoute struct {
	CalculationController *controller.CalculationController
}

func (cr *CalculationRoute) CalculationRoutes(rg *gin.RouterGroup) {
	rg.POST("/calculate-margin", cr.CalculationController.CalculateMargin)
}
