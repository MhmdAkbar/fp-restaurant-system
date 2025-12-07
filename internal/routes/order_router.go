package routes

import (
    ordercontroller "aplikasi_restoran/internal/controllers/order"
    orderdetailcontroller "aplikasi_restoran/internal/controllers/order_detail"
    "aplikasi_restoran/internal/middlewares"

    "github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.Engine, oc *ordercontroller.OrderController, odc *orderdetailcontroller.OrderDetailController) {

    // Public: tambah detail ke order
    order := r.Group("/orders")
    {
        order.POST("/:order_id/details", odc.AddDetail)
    }

    // Cashier: akses penuh ke order
    cashier := order.Group("")
    cashier.Use(middlewares.AuthMiddleware(), middlewares.Role("cashier"))
    {
        cashier.GET("/:order_id", oc.GetOrder)             // lihat order
        cashier.PUT("/details/:detail_id", odc.UpdateDetail) // update qty detail
        cashier.DELETE("/details/:detail_id", odc.DeleteDetail) // hapus detail
        // cashier.PUT("/:order_id", oc.UpdateOrder)          // update keseluruhan order
    }

    // Waiter: hanya update status
    waiter := order.Group("")
    waiter.Use(middlewares.AuthMiddleware(), middlewares.Role("waiter"))
    {
        waiter.PATCH("/:order_id/status", oc.UpdateStatus)
    }
}
