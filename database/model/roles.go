package model

type Roles struct {
	ID   int    `gorm:"primaryKey; autoIncrement" json:"id"`
	Role string `json:"role"`
}
