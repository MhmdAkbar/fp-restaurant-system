package orderdetailcontroller

import (
	"aplikasi_restoran/internal/dto"
	helpers "aplikasi_restoran/internal/helper"
	"aplikasi_restoran/internal/models"
	orderservice "aplikasi_restoran/internal/services/order"
	orderdetailservice "aplikasi_restoran/internal/services/order_detail"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderDetailController struct {
	service      orderdetailservice.OrderDetailService
	orderService orderservice.OrderService
}

func NewOrderDetailController(s orderdetailservice.OrderDetailService, os orderservice.OrderService) *OrderDetailController {
	return &OrderDetailController{
		service:      s,
		orderService: os,
	}
}

func (c OrderDetailController) AddDetail(ctx *gin.Context) {

	var input dto.AddOrderDetailRequest
	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	detail, err := c.service.AddDetail(input)
	//cek apakah order sudah dibayar
	order, _ := c.orderService.GetOrder(input.OrderId)
	if order.Status == models.OrderDone {
		helpers.ResponseError(ctx, http.StatusConflict, errors.New("this order has been paid, can't add detail"))
		return
	}
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	response := dto.OrderDetailResponse{
		ID:       detail.ID,
		OrderId:  detail.OrderId,
		MenuId:   detail.MenuId,
		Quantity: detail.Quantity,
		Price:    detail.Price,
		Total:    detail.Total,
		Menu: dto.SimpleMenu{
			Name:  detail.Menu.Name,
			Price: detail.Menu.Price,
		},
	}

	helpers.ResponseSuccess(ctx, http.StatusCreated, "Detail added", response)
}

func (c OrderDetailController) UpdateDetail(ctx *gin.Context) {
	// Ambil param ID dari URL
	detailID, ok := helpers.ParseID(ctx, "detail_id")
	if !ok {
		return
	}

	// Binding input qty
	var input dto.UpdateOrderDetailRequest
	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	updatedDetail, err := c.service.UpdateDetail(detailID, *input.Quantity)
	order, _ := c.orderService.GetOrder(updatedDetail.OrderId)
	if order.Status == models.OrderDone {
		helpers.ResponseError(ctx, http.StatusConflict, errors.New("this order has been paid, can't add detail"))
		return
	}
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	if updatedDetail == nil {
		_ = c.service.DeleteDetail(detailID)
		helpers.ResponseSuccess(ctx, http.StatusOK, "Order detail deleted", nil)
		return
	}

	// Berhasil update
	response := dto.OrderDetailResponse{
		ID:       updatedDetail.ID,
		OrderId:  updatedDetail.OrderId,
		MenuId:   updatedDetail.MenuId,
		Quantity: updatedDetail.Quantity,
		Price:    updatedDetail.Price,
		Total:    updatedDetail.Total,
		Menu: dto.SimpleMenu{
			Name:  updatedDetail.Menu.Name,
			Price: updatedDetail.Menu.Price,
		},
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Detail updated", response)
}

func (c OrderDetailController) DeleteDetail(ctx *gin.Context) {
	detailID, ok := helpers.ParseID(ctx, "order_id")
	if !ok {
		return
	}

	if err := c.service.DeleteDetail(detailID); err != nil {
		return
	}

	helpers.ResponseSuccess(ctx, http.StatusOK, "Detail deleted", nil)
}
