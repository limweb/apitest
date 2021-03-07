package models

type PasswordResets struct {
	ModelUpdatelog
	Email string `json:"email"  example:"" ` // PasswordResets email
	Token string `json:"token"  example:"" ` // PasswordResets token
}

func (m *PasswordResets) TableName() string {
	return "password_resets"
}
