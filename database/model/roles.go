package model

const (
	TableRolesName = "roles"
)

type Roles struct {
	ID   int `gorm:"primaryKey; autoIncrement"`
	Role string
}
