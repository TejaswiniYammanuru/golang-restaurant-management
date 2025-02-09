package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/models"
	"github.com/gin-gonic/gin"
)

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		orders, err := models.GetAllOrders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"orders": orders})
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("order_id")
		orderId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
			return
		}

		order, err := models.GetOrderByID(orderId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"order": order})
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		// Ensure table exists
		_, err := models.GetTableByID(order.TableID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Table not found"})
			return
		}

		// Set timestamps
		order.OrderDate = time.Now()
		order.CreatedAt = time.Now()
		order.UpdatedAt = time.Now()

		placedOrder, err := models.CreateOrder(&order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"order": placedOrder})
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("order_id")
		orderId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
			return
		}

		var updatedOrder models.Order
		if err := c.ShouldBindJSON(&updatedOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		order, err := models.GetOrderByID(orderId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		// Preserve old values if not updated
		if updatedOrder.OrderDate.IsZero() {
			updatedOrder.OrderDate = order.OrderDate
		}
		if updatedOrder.TableID == 0 {
			updatedOrder.TableID = order.TableID
		}
		if updatedOrder.UserID == 0 {
			updatedOrder.UserID = order.UserID
		}

		// Ensure table exists if changed
		if updatedOrder.TableID != order.TableID {
			_, err := models.GetTableByID(updatedOrder.TableID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Table not found"})
				return
			}
		}

		updatedOrder.ID = orderId
		updatedOrder.UpdatedAt = time.Now()

		err = models.UpdateOrder(&updatedOrder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully", "order": updatedOrder})
	}
}

func GetOrdersByUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userId, ok := userID.(int)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
			return
		}

		orders, err := models.GetOrdersByUserID(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"orders": orders})
	}
}
