package controller

import (
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
	serv service.Service
}

// função para receber os métodos da interface
func NewController(serv service.Service) Controller {
	return &controller{
		serv,
	}
}

func (ctl controller) PostUser(c *gin.Context) {
	userRequest := dto.UserRequest{}
	c.ShouldBindJSON(&userRequest)

	err := ctl.serv.UserService(userRequest)
	errors.IsEmptyError(err)

	c.JSON(http.StatusOK, "User created successfully!")
}
