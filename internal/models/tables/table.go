package tablemodels

import (
	"time"

	"gorm.io/gorm"
)

type Table struct {
	ID          uint   `gorm:"primaryKey"`
	TableCode string   `gorm:"not null"`
	QrCode      string `gorm:"size:255"`
	Status      bool   `gorm:"not null;default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
