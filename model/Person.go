package model

type Person struct {
	ID        uint   `json:”id”`        //Person ID
	FirstName string `json:”firstname”` //Person firstname
	LastName  string `json:”lastname”`  // Persion lastname
}
