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
		table.GET("/:id", tc.GetTable)                             // Get Single Table
		table.PATCH("/:id", tc.UpdateTable)                        // Update Full
		table.PATCH("/:id/status", tc.UpdateStatus)                // Update Hanya Status
		// table.DELETE("/:id", tc.DeleteTable)                    // Soft/Hard Delete (optional)
	}
}
