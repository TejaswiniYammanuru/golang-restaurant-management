package routes

import (
	"github.com/TejaswiniYammanuru/golang-restaurant-management/controllers"
	"github.com/TejaswiniYammanuru/golang-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(router *gin.Engine) {
	// Public routes (view menus)
	router.GET("/menus", controllers.GetMenus())
	router.GET("/menus/:menu_id", controllers.GetMenu())

	// Protect create, update, and delete routes with admin authentication
	adminRoutes := router.Group("/")
	adminRoutes.Use(middleware.AdminAuthentication())

	adminRoutes.POST("/menus", controllers.CreateMenu())
	adminRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
	adminRoutes.DELETE("/menus/:menu_id", controllers.DeleteMenu())
}
