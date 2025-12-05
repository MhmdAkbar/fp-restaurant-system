package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod string

const (
	PaymentCash  PaymentMethod = "cash"
	PaymentDebit PaymentMethod = "debit"
	PaymentQris  PaymentMethod = "qris"
)

type PaymentDone bool

const (
	Done    PaymentDone = true
	NotDone PaymentDone = false
)

type Payment struct {
	ID            uint          `gorm:"primaryKey"`
	OrderID       *uint         `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CashierID     uint          `gorm:"not null;index"`
	PaymentMethod PaymentMethod `gorm:"type:varchar(20);default:'cash'"`
	AmountPaid    float64       `gorm:"type:decimal(12,2);default:0"`
	ReferenceNo   string        `gorm:"type:varchar(50)"`
	PaidAt        *time.Time
	PaymentDone   bool `gorm:"default:false"`
	gorm.Model
}
