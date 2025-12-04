package routes

import (
    ordercontroller "aplikasi_restoran/internal/controllers/order"
    orderdetailcontroller "aplikasi_restoran/internal/controllers/order_detail"
    "aplikasi_restoran/internal/middlewares"

    "github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.Engine, oc *ordercontroller.OrderController, odc *orderdetailcontroller.OrderDetailController) {

    // =========================
    // Public/Internal Order Route
    // =========================
    order := r.Group("/orders")
    {
        // Add detail to order (waiter or system)
        order.POST("/:order_id/details", odc.AddDetail)
    }

    // =========================
    // Cashier Routes (Full Update)
    // =========================
    cashier := order.Group("")
    cashier.Use(middlewares.AuthMiddleware(), middlewares.Role("cashier"))
    {
        cashier.GET("/detail/:order_id", oc.GetOrder)   // lihat order

        // cashier.PUT("/:order_id", oc.UpdateOrder)       // cashier update isi pesanan
    }

    // =========================
    // Waiter Routes (Status Only)
    // =========================
    waiter := order.Group("")
    waiter.Use(middlewares.AuthMiddleware(), middlewares.Role("waiter"))
    {
        waiter.PATCH("/:order_id/status", oc.UpdateStatus) // waiter update status
    }
}
