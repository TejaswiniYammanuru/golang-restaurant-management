package routes

import (
	"github.com/TejaswiniYammanuru/golang-restaurant-management/controllers"
	"github.com/TejaswiniYammanuru/golang-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func TableRoutes(router *gin.Engine) {
	// Public routes (view tables)
	router.GET("/tables", controllers.GetTables())
	router.GET("/tables/:table_id", controllers.GetTable())

	// Protect create and update routes with admin authentication
	adminRoutes := router.Group("/")
	adminRoutes.Use(middleware.AdminAuthentication())

	adminRoutes.POST("/tables", controllers.CreateTable())
	adminRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())
}
