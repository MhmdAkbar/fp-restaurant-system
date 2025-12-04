package routes

import (
	ordercontroller "aplikasi_restoran/internal/controllers/order"
	tablecontroller "aplikasi_restoran/internal/controllers/table"
	"aplikasi_restoran/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func TableRouter(r *gin.Engine, tc *tablecontroller.TableController, or *ordercontroller.OrderController) {

	// ===================== PUBLIC (TANPA LOGIN) =====================
	public := r.Group("/tables")
	{
		public.POST("/:table_id/order", or.CreateOrder)
	}

	// ===================== ADMIN (HARUS LOGIN) =====================
	admin := r.Group("/tables")
	admin.Use(
		middlewares.AuthMiddleware(),
		middlewares.Role("admin"),
	)
	{
		admin.POST("", tc.AddTable)
		admin.GET("", tc.GetAll)
		admin.GET("/:table_id", tc.GetTable)
		admin.PATCH("/:table_id", tc.UpdateTable)
		admin.PATCH("/:table_id/status", tc.UpdateStatus)
	}
}
