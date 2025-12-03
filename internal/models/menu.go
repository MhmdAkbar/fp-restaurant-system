package models

import (
    "time"

    "gorm.io/gorm"
)

// Enum kategori menu
type MenuCategory string

const (
    Makanan MenuCategory = "makanan"
    Minuman MenuCategory = "minuman"
    Dessert MenuCategory = "dessert"
)

type Menu struct {
    ID          uint          `gorm:"primaryKey"`
    Name        string        `gorm:"type:varchar(150);not null"`
    Category    MenuCategory  `gorm:"type:enum('makanan','minuman','dessert');not null"`
    Price       float64       `gorm:"type:decimal(10,2);not null"`
    Description string        `gorm:"type:varchar(255)"`
    ImageURL    string        `gorm:"type:varchar(255)"`
    IsAvailable bool          `gorm:"default:true;not null"`
    Stock       int           `gorm:"default:0;not null"` // <-- baru: stok item

    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}
