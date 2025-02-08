package controllers

import (
	"strconv"
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/models"
	"github.com/gin-gonic/gin"
)

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoices, err := models.GetInvoices()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"invoices": invoices})

	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("invoice_id")
		invoiceId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Invalid invoice id"})
			return
		}
		invoice, err := models.GetInvoiceByID(invoiceId)
		if err != nil {
			c.JSON(404, gin.H{"error": "Invoice not found"})
			return

		}
		c.JSON(200, gin.H{"invoice": invoice})

	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invoice models.Invoice
		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		orderId:= invoice.OrderID

		_, err := models.GetOrderByID(orderId)
		if err != nil {
			c.JSON(404, gin.H{"error": "Order not found"})
			return
		}

		invoice.CreatedAt = time.Now()
		invoice.UpdatedAt = time.Now()
		err = models.CreateInvoice(&invoice)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Invoice created successfully", "invoice": invoice})

	}
}
func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("invoice_id")
		invoiceId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid invoice ID"})
			return
		}

		existingInvoice, err := models.GetInvoiceByID(invoiceId)
		if err != nil {
			c.JSON(404, gin.H{"error": "Invoice not found"})
			return
		}

		var updatedInvoice models.Invoice
		if err := c.ShouldBindJSON(&updatedInvoice); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if updatedInvoice.OrderID != 0 {

			orderId :=updatedInvoice.OrderID

			_, err := models.GetOrderByID(orderId)
			if err != nil {
				c.JSON(404, gin.H{"error": "Order not found"})
				return
			}
			existingInvoice.OrderID = updatedInvoice.OrderID;

		}

		

		if updatedInvoice.PaymentMethod != "" {
			existingInvoice.PaymentMethod = updatedInvoice.PaymentMethod
		}
		if updatedInvoice.PaymentStatus != "" {
			existingInvoice.PaymentStatus=updatedInvoice.PaymentStatus 
		}
		if !updatedInvoice.PaymentDueDate.IsZero() {
			existingInvoice.PaymentDueDate = updatedInvoice.PaymentDueDate
		}

		
		existingInvoice.UpdatedAt = time.Now()

		err = models.UpdateInvoice(existingInvoice)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Invoice updated successfully", "invoice": existingInvoice})
	}
}

// func DeleteInvoice() gin.HandlerFunc {
// 	return func(c *gin.Context) {
//         id := c.Param("invoice_id")
//         invoiceId, err := strconv.Atoi(id)
//         if err!= nil {
//             c.JSON(400, gin.H{"error": "Invalid invoice ID"})
//             return
//         }

//         err = models.DeleteInvoice(invoiceId)
//         if err!= nil {
//             c.JSON(500, gin.H{"error": err.Error()})
//             return
//         }

//         c.JSON(200, gin.H{"message": "Invoice deleted successfully"})
//     }
// }
