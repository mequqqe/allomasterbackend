package models

type RepairStatus struct {
	ID        uint `gorm:"primaryKey"`
	CompanyID uint
	Company   Company
	Name      string
	ColorHash string
}
