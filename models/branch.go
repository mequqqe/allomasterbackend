package models

type Branch struct {
	ID          uint `gorm:"primaryKey"`
	CompanyID   uint
	Company     Company
	BranchName  string
	Address     string
	PhoneNumber string
	Employees   []Employee
}
