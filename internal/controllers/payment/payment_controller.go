package paymentcontroller

import (
    "aplikasi_restoran/internal/dto"
    helpers "aplikasi_restoran/internal/helper"
    "aplikasi_restoran/internal/models"
    paymentservice "aplikasi_restoran/internal/services/payment"
    "errors"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type PaymentController struct {
    service paymentservice.PaymentService
}

func NewPaymentController(service paymentservice.PaymentService) *PaymentController {
    return &PaymentController{service: service}
}

func (c *PaymentController) CreatePayment(ctx *gin.Context) {
    // Ambil order_id dari URL
    orderParam := ctx.Param("order_id")
    orderID, err := strconv.ParseUint(orderParam, 10, 32)
    if err != nil {
        helpers.ResponseError(ctx, http.StatusBadRequest, errors.New("invalid order_id"))
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
    payment, err := c.service.CreatePayment(uint(orderID), cashierID, method)
    if err != nil {
        helpers.ResponseError(ctx, http.StatusInternalServerError, err)
        return
    }

    helpers.ResponseSuccess(ctx, http.StatusCreated, "success make payment", payment)
}
