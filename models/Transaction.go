package models

type Transactions struct {
	Transactions []Transactions `json:"transactions"`
}

type Transaction struct {
	ModelDefault
	TransactionForCreate
	ID uint `gorm:"primary_key"` // Transaction ID
}

type TransactionForCreate struct {
	TransactionForUpdate
}

type TransactionForUpdate struct {
	Total         float64 `json:"total"`          // Transaction Total
	Paid          float64 `json:"paid"`           // Transaction Paid
	Change        float64 `json:"change"`         // Transaction Change
	PaymentType   string  `json:"payment_type"`   // Transaction PaymentType
	PaymentDetail string  `json:"payment_detail"` // Transaction PaymentDetail
	OrderList     string  `json:"order_list"`     // Transaction OrderList
	StaffID       string  `json:"staff_id"`       // Transaction StaffID
}

func (m *Transaction) TableName() string {
	return "transactions"
}
