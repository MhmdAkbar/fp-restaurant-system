package routes

import (
	ordercontroller "aplikasi_restoran/internal/controllers/order"
	orderdetailcontroller "aplikasi_restoran/internal/controllers/order_detail"
	"aplikasi_restoran/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.Engine, oc *ordercontroller.OrderController, odc *orderdetailcontroller.OrderDetailController) {

	// Routes untuk admin & waiter (akses internal restoran)
	order := r.Group("/orders")
	{
		order.POST("/:id", oc.CreateOrder)
		order.POST("/:id/details", odc.AddDetail)	
	}

	cashier := order.Use(middlewares.AuthMiddleware(), middlewares.Role("cashier"))
	{
		cashier.GET("/detail/:id", oc.GetOrder)
	}
}
