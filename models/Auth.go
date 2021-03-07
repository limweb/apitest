package models

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

type Jwttoken struct {
	Token    string      `json:"token"  example:"" `    // Auth token
	Pubkey   string      `json:"pubkey"  example:"" `   // Auth public key
	Reftoken string      `json:"reftoken"  example:"" ` // Auth refresh token
	UserID   uint        `json:"userid" `               // Auth user ID
	Ulevel   uint        `json:"ulevel"  example:"1" `  // Auth ulevel
	Roles    interface{} `json:"roles" example:"" `     // Auth role
}
