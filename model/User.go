package model

import "time"

// https://gorm.io/docs/models.html

type User struct {
	ID        uint      `gorm:"primary_key"`                                   // User ID
	Username  string    `gorm:"unique" form:"username" binding:"required"`     // User username
	Password  string    `form:"password" binding:"required"`                   // User password
	Level     string    `gorm:"default:normal" example:"10" enums:"0,1,10,99"` // User level
	CreatedAt time.Time `swaggerignore:"true"`
}
