package models

import "time"

type Users struct {
	Users []Users `json:"users"`
}

type User struct {
	ModelDefault
	UserForCreate
	ID uint `json:"id" gorm:"primary_key" example:"1" ` // User id
}

type UserForCreate struct {
	UserForUpdate
	EmailVerifiedAt time.Time `json:"email_verified_at" example:"" `         // User emailverifiedat
	Level           uint      `json:"level" example:"99" `                   // User level
	RememberToken   string    `json:"token"  example:"aaaaa.bbbbbb.cccccc" ` // User token
}

type UserForUpdate struct {
	Name      string `json:"username" gorm:"unique" example:"usera" `       // User name
	Password  string `json:"password" example:"P@ssw0rd999" `               // User password
	Email     string `json:"email" gorm:"unique" example:"a@email.com" `    // User email
	Telephone string `json:"telephone" gorm:"unique" example:"0816477729" ` // User telephone
}

func (m *User) TableName() string {
	return "users"
}
