package models

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
}

type UserForUpdate struct {
	Username  string `json:"username" gorm:"unique" example:"usera" `          // User username
	Password  string `json:"password" example:"P@ssw0rd999" `                  // User password
	Telephone string `json:"telephone" gorm:"unique" example:"0816477729" `    // User telephone
	Level     uint   `json:"level" example:"99" `                              // User level
	Token     string `json:"token"  example:"afskdl;fjaslk;dfjaksl;dfjakl;s" ` // User token
}

func (m *User) TableName() string {
	return "users"
}
