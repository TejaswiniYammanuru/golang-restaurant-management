package routes

import (
	"github.com/TejaswiniYammanuru/golang-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/admin/login", controllers.AdminLogin)
	router.POST("/admin/register", controllers.AdminRegister)
}
