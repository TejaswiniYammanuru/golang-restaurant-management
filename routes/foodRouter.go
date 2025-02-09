package routes

import (
    "github.com/TejaswiniYammanuru/golang-restaurant-management/controllers"
    "github.com/TejaswiniYammanuru/golang-restaurant-management/middleware"
    "github.com/gin-gonic/gin"
)

func FoodRoutes(router *gin.Engine) {
    router.GET("/foods", controllers.GetFoods())
    router.GET("/foods/:food_id", controllers.GetFood())

    // Protect create, update, and delete routes with admin authentication
    adminRoutes := router.Group("/")
    adminRoutes.Use(middleware.AdminAuthentication())

    adminRoutes.POST("/foods", controllers.CreateFood())
    adminRoutes.PATCH("/foods/:food_id", controllers.UpdateFood())
    adminRoutes.DELETE("/foods/:food_id", controllers.DeleteFood())
}
