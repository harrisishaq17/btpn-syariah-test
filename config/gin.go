package config

import "github.com/gin-gonic/gin"

func NewGin() *gin.Engine {
	router := gin.Default()
	return router
}
