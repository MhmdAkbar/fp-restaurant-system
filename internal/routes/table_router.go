package routes

import (
	tablecontroller "aplikasi_restoran/internal/controllers/table"
	"aplikasi_restoran/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func TableRouter(r *gin.Engine, tc *tablecontroller.TableController) {
	table := r.Group("/tables")
	table.Use(middlewares.AuthMiddleware(),middlewares.Role("admin"))	
	{
		table.POST("/add-table", tc.AddTable)
		table.GET("/:id", tc.GetTable)
		table.PATCH("/:id/update-table", tc.UpdateTable)
		table.PATCH("/:id/update-status", tc.UpdateStatus)
		table.DELETE("/:id/delete", tc.DeleteTable)
	}
}