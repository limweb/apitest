package models

//database.AutoMigrate(&models.Test{}) //add to Db.go
type Tests struct {
	Tests []Tests `json:"tests"`
}

type Test struct {
	ModelDefault
	TestForCreate
	ID uint `json:"id" gorm:"primary_key" example:"1" ` // Test id
}

type TestForCreate struct {
	TestForUpdate
}

type TestForUpdate struct {
	Name string `json:"name" gorm:"index" example:"" ` // Test name
}

func (m *Test) TableName() string {
	return "tests"
}
