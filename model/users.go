package model

type Users struct {
	ID       int    `gorm:"primaryKey; autoIncrement" json:"id"`
	FullName string `gorm:"size:30" json:"full_name"`
	Email    string `gorm:"size:30; unique" json:"email"`
	Cpf      int64  `gorm:"unique" json:"cpf"`
	Password string `gorm:"size:15" json:"password"`
	RoleId   int    `json:"id_role"`
	Roles    Roles  `gorm:"foreignKey:RoleId"`
}
