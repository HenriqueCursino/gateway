package dto

type UserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Document string `json:"document" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   int    `json:"role_id" binding:"required"`
}

type UserCreate struct {
	FullName string
	Hash     string
	Email    string
	Document string
	Password string
	RoleId   int
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
