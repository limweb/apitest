package models

type Permissions struct {
	Permissions []Permissions `json:"permissions"`
}

type Permission struct {
	ModelDefault
	PermissionForCreate
}

type PermissionForCreate struct {
	PermissionForUpdate
}

type PermissionForUpdate struct {
	Name string `json:"name"  example:"create table" ` // Permission name
	Slug string `json:"slug"  example:"/aa" `          // Permission slug
	Desc string `json:"desc"  example:"คำอธิบาย" `     // Permission desc
}

func (m *Permission) TableName() string {
	return "permissions"
}
