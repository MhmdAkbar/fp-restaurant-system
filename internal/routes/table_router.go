package routes

import (
	tablecontroller "aplikasi_restoran/internal/controllers/table"
	"aplikasi_restoran/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func TableRouter(r *gin.Engine, tc *tablecontroller.TableController) {
	table := r.Group("/tables")
	table.Use(
		middlewares.AuthMiddleware(),
		middlewares.Role("admin"),
	)

	{
		table.POST("", tc.AddTable)                                // Create
		table.GET("", tc.GetAll)                                   // Get All Tables
		table.GET("/:table_id", tc.GetTable)                             // Get Single Table
		table.PATCH("/:table_id", tc.UpdateTable)                        // Update Full
		table.PATCH("/:table_id/status", tc.UpdateStatus)                // Update Hanya Status
		// table.DELETE("/:id", tc.DeleteTable)                    // Soft/Hard Delete (optional)
	}
}
