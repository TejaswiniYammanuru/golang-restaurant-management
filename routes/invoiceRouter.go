package routes

import (
	"github.com/TejaswiniYammanuru/golang-restaurant-management/controllers"
	"github.com/TejaswiniYammanuru/golang-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(router *gin.Engine) {
	// Public routes (view invoices)
	router.GET("/invoices", controllers.GetInvoices())
	router.GET("/invoices/:invoice_id", controllers.GetInvoice())

	// Protect create and update routes with admin authentication
	adminRoutes := router.Group("/")
	adminRoutes.Use(middleware.AdminAuthentication())

	adminRoutes.POST("/invoices", controllers.CreateInvoice())
	adminRoutes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
}
