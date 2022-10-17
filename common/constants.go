package common

const (
	Development                 = "development"
	CheckTableEmpty             = 1
	HeaderKey                   = "authorized"
	KeyHashToken                = "userHash"
	KeyExpToken                 = "exp"
	RemaningHoursToExpiredToken = 1440
	LenghtZero                  = 0
)

const (
	PermissionUserCreate = "user_create"
	PermissionGetUsers   = "user_read"
	PermissionUserDelete = "user_delete"
	PermissionUserUpdate = "user_update"

	PermissionRoleCreate = "role_create"
	PermissionRoleDelete = "role_delete"
	PermissionRoleUpdate = "role_update"
	PermissionGetRole    = "role_read"
)
