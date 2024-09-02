package models

import "time"

type Sale struct {
	ID             uint `gorm:"primaryKey"`
	BranchID       uint
	Branch         Branch `gorm:"foreignKey:BranchID"`
	CounterpartyID uint
	Counterparty   Counterparty `gorm:"foreignKey:CounterpartyID"`
	ProductID      uint
	Product        ProductDirectory `gorm:"foreignKey:ProductID"`
	ServiceID      uint
	Service        ServiceDirectory `gorm:"foreignKey:ServiceID"`
	Code           int
	Date           time.Time
	Price          float64
	TypePaymentID  uint
	TypePayment    TypePayment `gorm:"foreignKey:TypePaymentID"`
}
