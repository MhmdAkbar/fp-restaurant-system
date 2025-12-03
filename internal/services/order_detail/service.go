package orderdetailservice

import (
	"aplikasi_restoran/internal/dto"
	"aplikasi_restoran/internal/models"
)

type OrderDetailService interface {
	AddDetail(input dto.AddOrderDetailRequest) (*models.OrderDetail, error)
	// UpdateDetail(detailID uint, qty uint) (*models.OrderDetail, error)
	// DeleteDetail(detailID uint) error
	// GetOrderDetails(orderID uint) ([]models.OrderDetail, error)
}
