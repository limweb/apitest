package models


type RoleUsers struct {
	RoleUsers []RoleUsers `json:"role_users"`
}

type RoleUser struct {
	ModelDefault
	RoleUserForCreate
	ID uint `json:"id" gorm:"primary_key" example:"1" ` // RoleUser id
}

type RoleUserForCreate struct {
	RoleUserForUpdate
}

type RoleUserForUpdate struct {
	UserID uint `json:"user_id"  example:"1" ` // RoleUser userid
	RoleID uint `json:"role_id"  example:"1" ` // RoleUser roleid
}

func (m *RoleUser) TableName() string {
	return "role_user"
}
