package controller

import (
	"net/http"
	"sistem-pembiayaan/dto/calculation"
	"sistem-pembiayaan/service"

	"github.com/gin-gonic/gin"
)

type CalculationController struct {
	CalculationService service.CalculationService
}

func NewCalculationController(service service.CalculationService) *CalculationController {
	return &CalculationController{
		CalculationService: service,
	}
}

func (cc *CalculationController) CalculateMargin(c *gin.Context) {
	var req calculation.MarginCalculationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid"})
		return
	}

	data, status, err := cc.CalculationService.CalculateMargin(req.Amount)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, data)
}
