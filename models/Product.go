package models

type Products struct {
	Products []Products `json:"products"`
}

type Product struct {
	ModelDefault
	ProductForCreate
	ID string `json:"id" gorm:"primary_key" example:"1" ` // Product id
}

type ProductForCreate struct {
	ProductForUpdate
}

type ProductForUpdate struct {
	Name  string  `json:"name" example:"Mr. Aaa Bbbb" `      // Product name
	Stock int     `json:"stock"  example:"1" `               // Product stock
	Price float64 `json:"price"  example:"500.00" `          // Product price
	Image string  `json:"image" example:"/aaa/bbb/aaa.jpg" ` // Product image
}

func (m *Product) TableName() string {
	return "products"
}
