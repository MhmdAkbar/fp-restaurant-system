package ordercontroller

import (
	"aplikasi_restoran/internal/dto"
	helpers "aplikasi_restoran/internal/helper"
	"aplikasi_restoran/internal/models"
	orderservice "aplikasi_restoran/internal/services/order"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService orderservice.OrderService
}

func NewOrderController(os orderservice.OrderService) *OrderController {
	return &OrderController{orderService: os}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var input dto.CustomerCreateOrderRequest

	idTable, ok := helpers.ParseID(ctx, "table_id")
	if !ok {
		return
	}
	input.TableId = idTable

	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	order, err := c.orderService.CreateOrder(input)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusConflict, err)
		return
	}

	response := helpers.MapOrder(order)
	helpers.ResponseSuccess(ctx, http.StatusCreated, "Order created", response)
}

func (c OrderController) GetOrder(ctx *gin.Context) {
	id, ok := helpers.ParseID(ctx, "order_id")
	if !ok {
		return
	}

	order, err := c.orderService.GetOrder(id)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusNotFound, errors.New("order tidak ditemukan"))
		return
	}

	response := helpers.MapOrder(order)
	helpers.ResponseSuccess(ctx, http.StatusOK, "Success", response)
}

func (c *OrderController) UpdateStatus(ctx *gin.Context) {
	id, ok := helpers.ParseID(ctx, "order_id")
	if !ok {
		return
	}

	var req dto.UpdateOrderStatusRequest
	if !helpers.BindAndValidate(ctx, &req) {
		return
	}
	order, _ := c.orderService.GetOrder(id)
	if order.Status == models.OrderDone {
		helpers.ResponseError(ctx, http.StatusConflict, errors.New("this order has been paid, can't update status"))
		return
	}
	if err := c.orderService.UpdateStatus(id, req.Status); err != nil {
		helpers.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	updatedOrder, _ := c.orderService.GetOrder(id)
	response := helpers.MapOrder(updatedOrder)

	helpers.ResponseSuccess(ctx, http.StatusOK, "Status updated", response)
}
