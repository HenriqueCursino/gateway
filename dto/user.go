package dto

type UserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Cpf      string `json:"cpf"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}

type UserCreat struct {
	FullName string `gorm:"size:30"`
	Email    string `gorm:"size:30; unique" `
	Cpf      int64  `gorm:"unique"`
	Password string `gorm:"size:15"`
	RoleId   int
}
