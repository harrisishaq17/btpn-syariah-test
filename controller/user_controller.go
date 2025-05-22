package controller

import (
	"net/http"
	"sistem-pembiayaan/dto/user"
	"sistem-pembiayaan/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userService service.UserService
	validate    *validator.Validate
}

func NewUserController(userService service.UserService, validate *validator.Validate) *UserController {
	return &UserController{
		userService: userService,
		validate:    validate,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req user.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Format data tidak valid",
			"error":   err.Error(),
		})
		return
	}

	if err := uc.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validasi gagal",
			"error":   err.Error(),
		})
		return
	}

	status, err := uc.userService.CreateUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "User berhasil ditambahkan",
	})
}
