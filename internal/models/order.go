package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderReady   OrderStatus = "ready"
	OrderCooking OrderStatus = "cooking"
	OrderServed  OrderStatus = "served"
	OrderDone    OrderStatus = "done"
	OrderScanned OrderStatus = "scanned"
)

type OrderType string

const (
	OrderDineIn   OrderType = "dine_in"
	OrderTakeAway OrderType = "take_away"
)

type Order struct {
	ID            uint          `gorm:"primaryKey"`
	TableId       uint          `gorm:"not null;index"`
	Name          string        `gorm:"not null;index"`
	OrderType     OrderType     `gorm:"type:varchar(20);not null"`
	Status        OrderStatus   `gorm:"type:varchar(20);not null;default:'ready'"`
	Subtotal      float64       `gorm:"type:decimal(12,2);default:0"`
	Tax           float64       `gorm:"type:decimal(12,2);default:0"`
	ServiceCharge float64       `gorm:"type:decimal(12,2);default:0"`
	Discount      float64       `gorm:"type:decimal(12,2);default:0"`
	GrandTotal    float64       `gorm:"type:decimal(12,2);default:0"`
	Details       []OrderDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"details"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
