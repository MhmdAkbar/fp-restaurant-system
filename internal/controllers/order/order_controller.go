package ordercontroller

import (
	"aplikasi_restoran/internal/dto"
	helpers "aplikasi_restoran/internal/helper"
	orderservice "aplikasi_restoran/internal/services/order"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service orderservice.OrderService
}

func NewOrderController(s orderservice.OrderService) *OrderController {
	return &OrderController{service: s}
}
func (c OrderController) CreateOrder(ctx *gin.Context) {
	var input dto.CustomerCreateOrderRequest
    idTable, ok := helpers.ParseID(ctx,"tabel_id")
    if !ok {
        return
    }
    input.TableId = idTable
	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	order, err := c.service.CreateOrder(input)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, errors.New("ada yang salah"))
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusCreated, "Order created", order)
}

// func (c OrderController) GetOrCreateOrder(ctx *gin.Context) {
//     tableID, ok := helpers.ParseID(ctx)
//     if !ok {
//         return
//     }

//     // order, err := c.service.GetActiveOrderByTable(tableID)
//     // if err != nil {
//     //     // kalau tidak ada order aktif → buat baru
//     //     input := dto.CreateOrderRequest{
//     //         TableId: tableID,
//     //         Status:  "pending",
//     //     }
//     //     newOrder, err := c.service.CreateOrder(input)
//     //     if err != nil {
//     //         helpers.ResponseError(ctx, http.StatusBadRequest, err)
//     //         return
//     //     }
//     //     helpers.ResponseSuccess(ctx, http.StatusCreated, "Order created", newOrder)
//     //     return
//     // }

	
//     helpers.ResponseSuccess(ctx, http.StatusOK, "Order already active", order)
// }

func (c OrderController) GetOrder(ctx *gin.Context) {
	id, ok := helpers.ParseID(ctx,"order_id")
	if !ok {
		return
	}

	order, err := c.service.GetOrder(id)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusNotFound, err)
		return
	}
response := helpers.MapOrder(order)

helpers.ResponseSuccess(ctx, http.StatusOK, "Success", response)

}

// // func (c OrderController) GetActiveOrder(ctx *gin.Context) {
// // 	tableID, ok := helpers.ParseID(ctx)
// // 	if !ok {
// // 		return
// // 	}

// // 	order, err := c.service.GetActiveOrderByTable(tableID)
// // 	if err != nil {
// // 		helpers.ResponseError(ctx, http.StatusNotFound, err)
// // 		return
// 	}

// 	helpers.ResponseSuccess(ctx, http.StatusOK, "Success", order)
// }

// func (c OrderController) UpdateStatus(ctx *gin.Context) {
// 	id, ok := helpers.ParseID(ctx)
// 	if !ok {
// 		return
// 	}

// 	var input dto.UpdateOrderStatusRequest
// 	if !helpers.BindAndValidate(ctx, &input) {
// 		return
// 	}

// 	updated, err := c.service.UpdateOrderStatus(id, input.Status)
// 	if err != nil {
// 		helpers.ResponseError(ctx, http.StatusBadRequest, err)
// 		return
// 	}

// 	helpers.ResponseSuccess(ctx, http.StatusOK, "Status updated", updated)
// }

// func (c OrderController) DeleteOrder(ctx *gin.Context) {
// 	id, ok := helpers.ParseID(ctx)
// 	if !ok {
// 		return
// 	}

// 	err := c.service.DeleteOrder(id)
// 	if err != nil {
// 		helpers.ResponseError(ctx, http.StatusNotFound, err)
// 		return
// 	}

// 	helpers.ResponseSuccess(ctx, http.StatusOK, "Order deleted", nil)
// }
