package model

import "time"

type Product struct {
	ID        uint      `json:"id"; gorm:"primary_key"` // Product ID
	Name      string    `json:"name"`                   // Product Name
	Stock     int64     `json:"stock"`                  // Product Stock
	Price     float64   `json:"price"`                  // Product Price
	Image     string    `json:"image"`                  // Product Image
	CreatedAt time.Time `json:"time"; swaggerignore:"true"`
}
