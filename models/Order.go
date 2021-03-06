package models

type Orders struct {
	Orders []Orders `json:"orders"`
}

type Order struct {
	ModelDefault
	OrderForCreate
	Ccccc string `json:"ccccc" gorm:"index" example:"" ` // Order ccccc
}

type OrderForCreate struct {
	OrderForUpdate
	Bbbbb string `json:"bbbbb" gorm:"index" example:"" ` // Order bbbbb
}

type OrderForUpdate struct {
	Aaaaa string `json:"aaaaa" gorm:"index" example:"" ` // Order aaaaa
}

func (m *Order) TableName() string {
	return "orders"
}
