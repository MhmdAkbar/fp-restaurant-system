package routes

import (
	menucontroller "aplikasi_restoran/internal/controllers/menu"
	"aplikasi_restoran/internal/middlewares"

	"github.com/gin-gonic/gin"
)

// Route untuk fitur menu (hanya admin)
func MenuRouter(r *gin.Engine, mc *menucontroller.MenuController) {
	menu := r.Group("/menus")
	menu.Use(middlewares.AuthMiddleware(), middlewares.Role("admin")) // wajib login & admin

	{
		menu.POST("", mc.AddMenu)          // tambah menu
		menu.GET("", mc.GetAllMenu)        // list semua menu
		menu.GET("/:menu_id", mc.GetMenu)       // detail menu
		menu.PATCH("/:menu_id", mc.UpdateMenu)  // update menu
		menu.DELETE("/:menu_id", mc.DeleteMenu) // hapus menu
	}
}
