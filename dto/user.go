package dto

import "github.com/henriquecursino/gateway/database/model"

type UserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Document string `json:"document" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   int    `json:"role_id" binding:"required"`
}

type UserCreate struct {
	FullName string
	UserId   string
	Email    string
	Document string
	Password string
	RoleId   int
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToken struct {
	Email  string
	UserId string
}

type AllUsers struct {
	FullName string
	UserId   string
	Email    string
	Document string
	Password string
	Roles    model.Roles
}

type UpdateUserRole struct {
	Document string
	NewRole  int `json:"new_role_id"`
}

type UserDelete struct {
	UserId string `json:"user_id"`
}
