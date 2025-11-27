package usermodels

import (
	"time" // waktu

	"gorm.io/gorm" // gorm
)

type UserRole string

const (
	Admin      UserRole = "admin"       // admin
	Waiter     UserRole = "waiter"      // waiter
	Cashier    UserRole = "cashier"     // cashier
	SuperAdmin UserRole = "super_admin" // superadmin
)

type User struct {
	ID       uint     `gorm:"primaryKey"`                                          // primary key
	Name     string   `gorm:"type:varchar(255)"`                                   // nama user
	Email    string   `gorm:"unique;type:varchar(255)"`                            // email unik
	Password string   `gorm:"type:varchar(255)"`                                   // password hash
	Role     UserRole `gorm:"type:enum('admin','waiter','cashier','super_admin')"` // role user

	CreatedAt time.Time      // otomatis waktu buat
	UpdatedAt time.Time      // otomatis update
	DeletedAt gorm.DeletedAt `gorm:"index"` // soft delete
}
