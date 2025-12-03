package models

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	OrderId  uint    `gorm:"not null;index"`                    // foreign key ke orders
	MenuId   uint    `gorm:"not null;index"`                    // foreign key ke menus
	Quantity uint    `gorm:"not null"`                          // jumlah item
	Price    float64 `gorm:"type:decimal(10,2);not null"`       // snapshot harga per item
	Total    float64 `gorm:"type:decimal(10,2);not null"`       // price * quantity
	Menu     Menu    `gorm:"foreignKey:MenuId" json:"menu"`     // preload menu when needed
}
