package models

type Company struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Mail        string
	PhoneNumber string
	CompanyName string
	Password    string
	Branches    []Branch
	Employees   []Employee
	Devices     []DeviceDirectory
	Roles       []Role
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
