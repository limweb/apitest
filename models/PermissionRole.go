package models

type PermissionRoles struct {
	PermissionRoles []PermissionRoles `json:"permissionroles"`
}

type PermissionRole struct {
	ModelDefault
	PermissionRoleForCreate
	ID uint `json:"id" gorm:"primary_key" example:"1" ` // PermissionRole id
}

type PermissionRoleForCreate struct {
	PermissionRoleForUpdate
}

type PermissionRoleForUpdate struct {
	PermissionID uint `json:"permission_id" example:"1" ` // PermissionRole permissionid
	RoleID       uint `json:"role_id"  example:"1" `      // PermissionRole roleid
}

func (m *PermissionRole) TableName() string {
	return "permission_role"
}
