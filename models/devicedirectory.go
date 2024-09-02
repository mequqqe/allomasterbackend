package models

type DeviceDirectory struct {
	ID        uint `gorm:"primaryKey"`
	CompanyID uint
	Company   Company
	Group     string
	Brand     string
	Model     string
}
