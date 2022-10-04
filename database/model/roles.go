package model

type Roles struct {
	ID   int `gorm:"primaryKey; autoIncrement"`
	Role string
}
