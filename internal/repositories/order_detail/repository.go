package orderdetailrepository

import (
	"aplikasi_restoran/internal/models"
)

type OrderDetailRepository interface {
	Create(detail *models.OrderDetail) (*models.OrderDetail, error) // buat detail order
	// GetByOrderID(orderID uint) ([]models.OrderDetail, error) // ambil detail per order dan dikembalikan dalam bentuk slice
	// GetByID(id uint) (*models.OrderDetail, error)
	// DeleteByOrderID(orderID uint) error // hapus seluruh detail berdasar order
	// Update(detail *models.OrderDetail) error
	// Delete(detailID uint) error
}
