package models

type Persons struct {
	Persons []Persons `json:"persons"`
}

type Person struct {
	ID uint `json:"id" gorm:"primary_key" example:"1" ` // Person id
	ModelDefault
	PersonForCreate
}

type PersonForCreate struct {
	PersonForUpdate
}

type PersonForUpdate struct {
	FirstName string `json:"firstname" example:"Aaaa" ` // Person firstname
	LastName  string `json:"lastname"  example:"Lim" `  // Person lastname
}

func (m *Person) TableName() string {
	return "persons"
}
