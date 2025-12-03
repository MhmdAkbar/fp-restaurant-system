package dto

// REQUEST DTO =========================

type UpdateOrderDetailRequest struct {
	Quantity uint `json:"quantity" binding:"required,min=1"`
}

type AddOrderDetailRequest struct {
	OrderId  uint `json:"order_id" binding:"required"`
	MenuId   uint `json:"menu_id" binding:"required"`
	Quantity uint `json:"quantity" binding:"required,min=1"`
}

// RESPONSE DTO =========================
type SimpleMenu struct {
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

type OrderDetailResponse struct {
    ID       uint       `json:"id"`
    OrderId  uint       `json:"order_id"`
    MenuId   uint       `json:"menu_id"`
    Quantity uint       `json:"quantity"`
    Price    float64    `json:"price"`
    Total    float64    `json:"total"`
    Menu     SimpleMenu `json:"menu"`
}

