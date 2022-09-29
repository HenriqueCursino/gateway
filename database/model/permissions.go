package model

type Permissions struct {
	ID         int    `gorm:"primaryKey; autoIncrement"`
	Permission string `gorm:"size:30"`
}
