package dto

type RoleUser struct {
	Role          string `json:"role"`
	PermissionsId []int  `json:"permissions_id"`
}

type PermissionRole struct {
	RoleId       int
	PermissionId int
}
