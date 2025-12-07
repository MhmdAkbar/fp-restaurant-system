package paymentcontroller

import (
	"aplikasi_restoran/internal/dto"
	helpers "aplikasi_restoran/internal/helper"
	"aplikasi_restoran/internal/models"
	orderservice "aplikasi_restoran/internal/services/order"
	paymentservice "aplikasi_restoran/internal/services/payment"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	service paymentservice.PaymentService
    order orderservice.OrderService
}

func NewPaymentController(service paymentservice.PaymentService, order orderservice.OrderService) *PaymentController {
	return &PaymentController{
        service: service,
        order: order,
    
    }
}

func (c *PaymentController) CreatePayment(ctx *gin.Context) {
	// Ambil order_id dari URL
	orderID, ok := helpers.ParseID(ctx, "order_id")
	if !ok {
		return
	}
    
    order , err :=c.order.GetOrder(orderID)
    if err != nil {
        helpers.ResponseError(ctx,  http.StatusNotFound, errors.New("order is not found"))
        return
    }
    if order.Status == models.OrderDone {
        helpers.ResponseError(ctx, http.StatusConflict, errors.New("this order has been paid"))
        return
    }
	// Ambil cashier id dari JWT (set di middleware)
	cashierVal, ok := ctx.Get("user_id")
	if !ok {
		helpers.ResponseError(ctx, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	cashierID, ok := cashierVal.(uint)
	if !ok {
		// kadang middleware menyimpan int64 — adaptasi jika perlu
		if v64, ok2 := cashierVal.(uint64); ok2 {
			cashierID = uint(v64)
		} else {
			helpers.ResponseError(ctx, http.StatusUnauthorized, errors.New("invalid user id in token"))
			return
		}
	}

	// Bind body (only method)
	var req dto.CreatePaymentRequest
	if !helpers.BindAndValidate(ctx, &req) {
		helpers.ResponseError(ctx, http.StatusBadRequest, errors.New("format JSON salah"))
		return
	}

	method := models.PaymentMethod(req.Method)

	// Panggil service — amount otomatis dari order.GrandTotal
	payment, err := c.service.CreatePayment(orderID, cashierID, method)
	if err != nil {
        fmt.Println("terjadai error di baris 70 payment controller")
		helpers.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusCreated, "success make payment", payment)
}
