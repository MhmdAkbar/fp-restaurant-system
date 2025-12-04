package ordercontroller

import (
	"aplikasi_restoran/internal/dto"
	helpers "aplikasi_restoran/internal/helper"
	orderdetailservice "aplikasi_restoran/internal/services/order_detail"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderDetailController struct {
	service orderdetailservice.OrderDetailService
}

func NewOrderDetailController(s orderdetailservice.OrderDetailService) *OrderDetailController {
	return &OrderDetailController{service: s}
}

func (c OrderDetailController) AddDetail(ctx *gin.Context) {

	var input dto.AddOrderDetailRequest
	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	detail, err := c.service.AddDetail(input)
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
