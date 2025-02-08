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
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"orders": orders})

	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("order_id")
		orderId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Invalid order id"})
			return
		}
		order, err := models.GetOrderByID(orderId)
		if err != nil {
			c.JSON(404, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(200, gin.H{"order": order})

	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		if err := c.ShouldBindJSON(&order); err!= nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
		order.OrderDate = time.Now()
		//check if that table with id TableID exists
		_, err := models.GetTableByID(order.TableID)
        if err!= nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Table not found"})
            return
        }
        placedorder,err := models.CreateOrder(&order)
        if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"order": placedorder})
		

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

		
		updatedOrder.ID = orderId

		
		if updatedOrder.OrderDate.IsZero() {
			updatedOrder.OrderDate = order.OrderDate
		}
		if updatedOrder.TableID == 0 {
			updatedOrder.TableID = order.TableID
		}

		
		if updatedOrder.TableID != 0 {
			_, err := models.GetTableByID(updatedOrder.TableID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Table not found"})
				return
			}
		}

		updatedOrder.UpdatedAt = time.Now()
		err = models.UpdateOrder(&updatedOrder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully", "order": updatedOrder})
	}
}


