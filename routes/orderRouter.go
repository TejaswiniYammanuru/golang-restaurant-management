package routes

import (
	"github.com/TejaswiniYammanuru/golang-restaurant-management/controllers"
	"github.com/TejaswiniYammanuru/golang-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	// Public routes (view orders)
	
	router.GET("/orders/:order_id", controllers.GetOrder())
	router.GET("/orders", controllers.GetOrdersByUserID())
	router.POST("/orders", controllers.CreateOrder())
	
	// Protect create and update routes with admin authentication
	adminRoutes := router.Group("/")
	adminRoutes.Use(middleware.AdminAuthentication())
	adminRoutes.GET("/allorders", controllers.GetOrders())
	
	adminRoutes.PATCH("/orders/:order_id", controllers.UpdateOrder())
}
