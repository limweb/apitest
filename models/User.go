package models

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Users struct {
	Users []Users `json:"users"`
}

type User struct {
	ModelDefault
	UserForCreate
	Roles []*Role `gorm:"many2many:role_users"`
}

type UserForCreate struct {
	UserForUpdate
	EmailVerifiedAt sql.NullTime `json:"email_verified_at" swaggerignore:"true" sql:"default:null" example:"" ` // User emailverifiedat
	Level           uint         `json:"level" example:"99" `                                                   // User level
	RememberToken   string       `json:"reftoken"  example:"aaaaa.bbbbbb.cccccc" `                              // User token
}

type UserForUpdate struct {
	Name      string `json:"name" example:"Mr. A " `                        // User password
	Username  string `json:"username" gorm:"unique" example:"usera" `       // User name
	Password  string `model:"hide" json:"password" example:"P@ssw0rd999" `  // User password
	Email     string `json:"email" gorm:"unique" example:"a@email.com" `    // User email
	Telephone string `json:"telephone" gorm:"unique" example:"0816477729" ` // User telephone
}

func (m *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.RememberToken = uuid.NewV4().String()
	return
}

// func (u *User) AfterCreate(tx *gorm.DB) (err error) {
// 	u.Password = "---hide---"
// 	return
// }
