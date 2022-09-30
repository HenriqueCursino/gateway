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
	}

	ctx.JSON(http.StatusOK, "User created successfully!")
}
