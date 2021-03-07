package models

type Roles struct {
	Roles []Roles `json:"roles"`
}

type Role struct {
	ModelDefault
	RoleForCreate
	ID uint `json:"id" gorm:"primary_key" example:"1" ` // Role id
}

type RoleForCreate struct {
	RoleForUpdate
}

type RoleForUpdate struct {
	Name string `json:"name"  example:"Admin" ` // Role name
}

func (m *Role) TableName() string {
	return "roles"
}
