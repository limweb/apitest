package models

type Premissions struct {
	Premissions []Premissions `json:"premissions"`
}

type Premission struct {
	ModelDefault
	PremissionForCreate
	ID uint `json:"id" gorm:"primary_key" example:"1" ` // Premission id
}

type PremissionForCreate struct {
	PremissionForUpdate
}

type PremissionForUpdate struct {
	Name string `json:"name"  example:"create table" ` // Premission name
	Slug string `json:"slug"  example:"/aa" `          // Premission slug
	Desc string `json:"desc"  example:"คำอธิบาย" `     // Premission desc
}

func (m *Premission) TableName() string {
	return "premissions"
}
