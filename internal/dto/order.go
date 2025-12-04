package dto

import "aplikasi_restoran/internal/models"

// REQUEST DTO =========================

type CustomerCreateOrderRequest struct {
	TableId   uint                `json:"table_id" binding:"required"`
	Name      string              `json:"name" binding:"required"`
	OrderType models.OrderType    `json:"order_type" binding:"required,oneof=dine_in take_away"`
	Status    models.OrderStatus  `json:"status" binding:"required,oneof=scanned cooking ready served done"`
	Details   []CreateOrderDetail `json:"details" binding:"required,dive"`
}

type CreateOrderDetail struct {
	MenuId   uint `json:"menu_id" binding:"required"`
	Quantity uint `json:"quantity" binding:"required,min=1"`
}

type UpdateOrderStatusRequest struct {
	Status models.OrderStatus `json:"status" binding:"required,oneof=cooking ready served done"`
}

// RESPONSE DTO =========================

type OrderResponse struct {
	ID            uint                  `json:"id"`
	TableId       uint                  `json:"table_id"`
	Name          string                `json:"name" binding:"required"`
	OrderType     models.OrderType      `json:"order_type"`
	Status        models.OrderStatus    `json:"status"`
	Subtotal      float64               `json:"subtotal"`
	Tax           float64               `json:"tax"`
	ServiceCharge float64               `json:"service_charge"`
	Discount      float64               `json:"discount"`
	GrandTotal    float64               `json:"grand_total"`
	OrderDetails  []OrderDetailResponse `json:"order_details"`
}

type ActiveOrderResponse struct {
	OrderResponse
	IsActive bool `json:"is_active"`
}
