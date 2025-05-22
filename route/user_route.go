package route

import (
	"sistem-pembiayaan/controller"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	UserController *controller.UserController
}

func NewUserRoute(userController *controller.UserController) *UserRoute {
	return &UserRoute{
		UserController: userController,
	}
}

func (ur *UserRoute) RegisterRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	{
		user.POST("/", ur.UserController.CreateUser)
	}
}
