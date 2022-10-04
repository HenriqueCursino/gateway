package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/common/errors"
	"github.com/henriquecursino/gateway/dto"
	"github.com/henriquecursino/gateway/service"
)

type Controller interface {
	PostUser(c *gin.Context)
	Login(ctx *gin.Context)
}

type controller struct {
	service service.Service
}

// NewController receive methods about core user
func NewController(service service.Service) Controller {
	return &controller{
		service,
	}
}

func (c *controller) PostUser(ctx *gin.Context) {
	userRequest := dto.UserRequest{}
	errBindJSON := ctx.ShouldBindJSON(&userRequest)
	if errBindJSON != nil {
		log.Fatal("Failed to bind JSON! - ", errBindJSON)
	}

	err := c.service.UserService(userRequest)
	if !errors.IsEmptyError(err) {
		ctx.JSON(http.StatusBadRequest, "Failed to create user!")
		return
	}

	ctx.JSON(http.StatusOK, "User created successfully!")
}

func (c *controller) Login(ctx *gin.Context) {
	loginRequest := dto.UserLogin{}
	errBindJSON := ctx.ShouldBindJSON(&loginRequest)
	if errBindJSON != nil {
		log.Fatal("Failed to bind JSON! - ", errBindJSON)
	}

	err := c.service.LoginService(loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to login!")
		return
	}

	token, errorToken := c.service.CreateJWT(&loginRequest)
	if errorToken != nil {
		ctx.JSON(http.StatusInternalServerError, errorToken)
	}

	ctx.JSON(http.StatusOK, token)
}
