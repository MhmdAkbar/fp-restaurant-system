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

func (c OrderDetailController) UpdateDetail(ctx *gin.Context) {
	// Ambil param ID dari URL
	detailID, ok := helpers.ParseID(ctx, "order_id")
	if !ok {
		return
	}

	// Binding input qty
	var input dto.UpdateOrderDetailRequest
	if !helpers.BindAndValidate(ctx, &input) {
		return
	}

	// Update via service
	updatedDetail, err := c.service.UpdateDetail(detailID, input.Quantity)
	if err != nil {
		helpers.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	// Jika qty 0 -> detail telah dihapus
	if updatedDetail == nil {
		helpers.ResponseSuccess(ctx, http.StatusOK, "Detail deleted", nil)
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
