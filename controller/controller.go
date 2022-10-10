package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/common"
	"github.com/henriquecursino/gateway/common/errors"
	"github.com/henriquecursino/gateway/dto"
	"github.com/henriquecursino/gateway/middleware"
	"github.com/henriquecursino/gateway/service"
)

type Controller interface {
	PostUser(c *gin.Context)
	Login(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
}

type controller struct {
	service    service.Service
	middleware middleware.Middleware
}

// NewController receive methods about core user
func NewController(service service.Service, middleware middleware.Middleware) Controller {
	return &controller{
		service,
		middleware,
	}
}

func (c *controller) PostUser(ctx *gin.Context) {
	if c.middleware.CheckPermission(ctx, common.PermissionUserCreate) {
		userRequest := dto.UserRequest{}
		if errBindJSON := ctx.ShouldBindJSON(&userRequest); errBindJSON != nil {
			log.Fatal("Failed to bind JSON! - ", errBindJSON)
		}

		err := c.service.UserService(userRequest)
		if !errors.IsEmptyError(err) {
			ctx.JSON(http.StatusBadRequest, "Failed to create user!")
			return
		}

		ctx.JSON(http.StatusOK, "User created successfully!")
		return
	}
	ctx.JSON(http.StatusBadRequest, "Doesn't have permission to create a new user!")
}

func (c *controller) Login(ctx *gin.Context) {
	loginRequest := dto.UserLogin{}
	errBindJSON := ctx.ShouldBindJSON(&loginRequest)
	if errBindJSON != nil {
		log.Fatal("Failed to bind JSON! - ", errBindJSON)
	}

	user, err := c.service.LoginService(loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to login!")
		return
	}

	token, errorToken := c.service.CreateJWT(user)
	if errorToken != nil {
		ctx.JSON(http.StatusInternalServerError, errorToken)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (c *controller) GetAllUsers(ctx *gin.Context) {
	if c.middleware.CheckPermission(ctx, common.PermissionGetUsers) {
		allUsers, err := c.service.GetAllUsersService()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Failed to get all users")
		}
		ctx.JSON(http.StatusOK, allUsers)
		return
	}
	ctx.JSON(http.StatusBadRequest, "Doesn't have permission to get all users!")
}
