package models

import (
	"time"

	"gorm.io/gorm"
)
type Table struct {
    ID        uint           `gorm:"primaryKey"`
    TableCode string         `gorm:"size:50;not null;uniqueIndex"`
    QrCode    string         `gorm:"size:255"`
    Status    bool           `gorm:"not null;default:false"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

