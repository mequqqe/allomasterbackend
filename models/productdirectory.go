package models

type ProductDirectory struct {
	ID                 uint `gorm:"primaryKey"`
	CompanyID          uint
	Company            Company
	ProductName        string
	ProductDescription string
	ProductPrice       string
	ProductImage       string
	Guarantee          string
	Code               int
	OptBalance         int
	PurchasePrice      int
	RetailPrice        int
}
