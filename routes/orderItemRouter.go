package routes

import (
	"github.com/TejaswiniYammanuru/golang-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(router *gin.Engine) {
	// Public routes (view order items)
	router.GET("/orderItems", controllers.GetOrderItems())
	router.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	router.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())

	// Protected routes (create and update order items)
	router.POST("/orderItems", controllers.CreateOrderItem())
	router.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
