package models

type TypePayment struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CompanyID uint
	Company   Company
}
