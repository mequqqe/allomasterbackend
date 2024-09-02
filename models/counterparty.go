package models

type Counterparty struct {
	ID           uint `gorm:"primaryKey"`
	CompanyID    uint
	Company      Company
	Organization bool
	Name         string
	Address      string
	Phone        string
	Email        string
	Comment      string
	Description  string
	BankDetails  string
}
