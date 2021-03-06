package models

type Books struct {
	Books []Books `json:"books"`
}

type Book struct {
	ModelDefault
	BookForCreate
	ID uint `json:"id" gorm:"primary_key" example:"1" ` // Book id
}

type BookForCreate struct {
	BookForUpdate
}

type BookForUpdate struct {
	Name     string `json:"name" example:"Mr. Aaaa Bbbbb" `  // Book name
	Author   string `json:"author" example:"Mr. Cccc Dddd" ` // Book author
	Category string `json:"category" example:"Sea" `         // Book category
}

func (m *Book) TableName() string {
	return "books"
}
