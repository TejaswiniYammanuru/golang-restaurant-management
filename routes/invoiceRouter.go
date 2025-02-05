package routes

import (
	controller "github.com/TejaswiniYammanuru/golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)



func InvoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id",controller.GetInvoice())
	incomingRoutes.POST("/invoices",controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id",controller.UpdateInvoice())

}