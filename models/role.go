package models

type Role struct {
	ID          uint `gorm:"primaryKey"`
	CompanyID   uint
	Company     Company
	RoleName    string
	Permissions string
}
