package models

type OrderDetail struct {
	ID uint `gorm:"primaryKey"` // id utama order detail
	OrderId uint `gorm:"not null"` // id ke order induk
	MenuId uint `gorm:"not null"` // id menu yang dipesan
	Quantity uint `gorm:"not null"` // jumlah item
	Price float64 `gorm:"type:decimal(10,2);not null"` // snapshoot harga satuan (agar meskipun harga berubah detail tetap aman)
	Total float64 `gorm:"type:decimal(10,2);not null"` // total harga untuk detail ini (price * quantity)
}