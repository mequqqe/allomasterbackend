package models

type ServiceDirectory struct {
	ID          uint `gorm:"primaryKey"`
	CompanyID   uint
	Company     Company
	ServiceName string
	Code        int
	Guarantee   int
	Price       int
}
