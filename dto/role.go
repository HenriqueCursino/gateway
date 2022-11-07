package dto

type RoleUser struct {
	Role          string `json:"role"`
	PermissionsId []int  `json:"permissions_id"`
}

type PermissionRole struct {
	RoleId       int
	PermissionId *int
}

type AllRoles struct {
	Role        string
	Permissions []Permissions
}

type Permissions struct {
	Permission string
}

type RoleDelete struct {
	RoleId int `json:"role_id"`
}
