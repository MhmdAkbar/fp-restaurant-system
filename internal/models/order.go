package models

import "gorm.io/gorm"

type OrderStatus string // tipe status order

const (
	OrderCooking OrderStatus = "cooking" //order sudah masuk dan sedang dimasak
	OrderReady   OrderStatus = "ready"   //order sudah siap dan siap diantar
	OrderServed  OrderStatus = "served"  //order sudah dihidangkan
	OrderDone    OrderStatus = "done"    // order sudah selesai dan sudah dibayar
)

type OrderType string

const (
	OrderDineIn   OrderType = "dine_in"   // makan ditempat
	OrderTakeAway OrderType = "take_away" // dibawa pulang
)

type Order struct {
	ID            uint          `gorm:"primaryKey"`                                   // id utama
	TableId       uint          `gorm:"not null"`                                     // id meja atau id table
	UserId        uint          `gorm:"not null"`                                     // id waiter yang melayani pesanan ini
	OrderType     OrderType     `gorm:"type:enum('dine_in','take_away');not null"`    // tipe pesanan apakah makan ditempat atau bungkus
	Status        OrderStatus   `gorm:"type:enum('cooking','ready','served','done')"` // status pesanan
	Subtotal      float64       `gorm:"default:0"`                                    // total pesanan snapshot dari order details
	Tax           float64       `gorm:"default:0"`                                    // pajak (default 0)
	ServiceCharge float64       `gorm:"default:0"`                                    // biaya service (default 0)
	Discount      float64       `gorm:"default:0"`                                    // diskon default 0
	GrandTotal    float64       `gorm:"default:0"`                                    // total akhir setelah perhitungn pajak dll
	OrderDetails  []OrderDetail `gorm:"foreignKey:OrderId"`                           //relasi ke table order detail
	gorm.Model                  // created, updated, soft delete
}
