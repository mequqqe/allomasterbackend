package models

import "time"

type Repair struct {
	ID                uint `gorm:"primaryKey"`
	BranchID          uint
	Branch            Branch
	Date              time.Time
	CounterpartyID    uint
	Counterparty      Counterparty
	Phone             string
	DeviceID          uint
	Device            DeviceDirectory
	Code              int
	Completeness      string
	ReasonForContract string
	EmployeeID        uint
	Employee          Employee
	Price             int
	Comment           string
}
