package helpers

import (
	"aplikasi_restoran/internal/dto"
	"aplikasi_restoran/internal/models"
)

func MapOrderDetails(details []models.OrderDetail) []dto.OrderDetailResponse {
	res := make([]dto.OrderDetailResponse, 0, len(details))

	for _, d := range details {

		// Menu ringkas — cek ID bukan nol
		menu := dto.SimpleMenu{}
		if d.Menu.ID != 0 {
			menu.Name = d.Menu.Name
			menu.Price = d.Menu.Price
		}

		res = append(res, dto.OrderDetailResponse{
			ID:       d.ID,
			OrderId:  d.OrderId,
			Name:     d.Name,
			MenuId:   d.MenuId,
			Quantity: d.Quantity,
			Price:    d.Price,
			Total:    d.Total,
			Menu:     menu,
		})
	}

	return res
}

func MapOrder(order *models.Order) dto.OrderResponse {
	return dto.OrderResponse{
		ID:            order.ID,
		TableId:       order.TableId,
		Name: order.Name,
		OrderType:     order.OrderType,
		Status:        order.Status,
		Subtotal:      order.Subtotal,
		Tax:           order.Tax,
		ServiceCharge: order.ServiceCharge,
		Discount:      order.Discount,
		GrandTotal:    order.GrandTotal,
		OrderDetails:  MapOrderDetails(order.Details),
	}
}

// func IsDone (tableId uint) bool {
//     table :=
// }
