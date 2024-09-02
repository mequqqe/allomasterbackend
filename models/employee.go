package models

type Employee struct {
	ID          uint `gorm:"primaryKey"`
	BranchID    uint
	Branch      Branch
	Name        string
	Mail        string
	PhoneNumber string
	Password    string
	RoleID      uint
	Role        Role
	CompanyID   uint
	Company     Company
}
