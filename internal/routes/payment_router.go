package routes

import (
	paymentcontroller "aplikasi_restoran/internal/controllers/payment"
	"aplikasi_restoran/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func PaymentRouter(r *gin.Engine, pc *paymentcontroller.PaymentController) {
	cashier := r.Group("/payment")
	cashier.Use(middlewares.AuthMiddleware(), middlewares.Role("cashier"))
	{
		cashier.POST("/:order_id", pc.CreatePayment)
	}
}