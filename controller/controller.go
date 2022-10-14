package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/gateway/common/errors"
	"github.com/henriquecursino/gateway/dto"
	"github.com/henriquecursino/gateway/middleware"
	"github.com/henriquecursino/gateway/service"
)

type Controller interface {
	PostUser(c *gin.Context)
	Login(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	UpdateUserRole(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	PostRole(ctx *gin.Context)
	GetAllRoles(ctx *gin.Context)
}

type controller struct {
	service    service.Service
	middleware middleware.Middleware
}

func NewController(service service.Service, middleware middleware.Middleware) Controller {
	return &controller{
		service,
		middleware,
	}
}

func (c *controller) PostUser(ctx *gin.Context) {
	var userRequest dto.UserRequest
	if errBindJSON := ctx.ShouldBindJSON(&userRequest); errBindJSON != nil {
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
	var loginRequest dto.UserLogin
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
	allUsers, err := c.service.GetAllUsersService()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to get all users")
		return
	}
	ctx.JSON(http.StatusOK, allUsers)
}

func (c *controller) DeleteUser(ctx *gin.Context) {
	var deleteRequest dto.UserDelete
	errBindJSON := ctx.ShouldBindJSON(&deleteRequest)
	if errBindJSON != nil {
		log.Fatal("Failed to bind JSON! - ", errBindJSON)
	}
	err := c.service.DeleteUserService(deleteRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to get all users")
		return
	}
	ctx.JSON(http.StatusOK, "User deleted successfully!")
}

func (c *controller) UpdateUserRole(ctx *gin.Context) {
	document := ctx.Param("doc")
	updateUser := dto.UpdateUserRole{
		Document: document,
	}
	errBindJSON := ctx.ShouldBindJSON(&updateUser)
	if errBindJSON != nil {
		log.Fatal("Failed to bind JSON! - ", errBindJSON)
	}

	err := c.service.UpdateUserRole(updateUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to update user role!")
		return
	}
	ctx.JSON(http.StatusOK, "User role successfully updated!")
}

func (c *controller) PostRole(ctx *gin.Context) {
	var createRole dto.RoleUser
	errBindJSON := ctx.ShouldBindJSON(&createRole)
	if errBindJSON != nil {
		log.Fatal("Failed to bind JSON! - ", errBindJSON)
	}

	err := c.service.CreateRole(createRole)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to createS new role!")
		return
	}
	ctx.JSON(http.StatusOK, "User role successfully created!")
}

func (c *controller) GetAllRoles(ctx *gin.Context) {
	allRoles, err := c.service.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to get all roles!")
	}
	ctx.JSON(http.StatusOK, allRoles)
}
