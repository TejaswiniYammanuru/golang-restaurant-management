package controllers

import (
	"strconv"
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/models"
	"github.com/gin-gonic/gin"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItems, err := models.GetOrderItems()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"orderItems": orderItems})

	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("orderItem_id")
		orderItemId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Invalid order item id"})
			return
		}
		orderItem, err := models.GetOrderItem(orderItemId)
		if err != nil {
			c.JSON(404, gin.H{"error": "Order item not found"})
			return
		}
		c.JSON(200, gin.H{"orderItem": orderItem})

	}
}

func CreateOrderItem() gin.HandlerFunc {

	return func(c *gin.Context) {
		var orderItem models.OrderItem
		if err := c.ShouldBindJSON(&orderItem); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		//check if that Food with that FoodID exists and OrderID too
		_, err := models.GetFoodByID(orderItem.FoodID)
		if err != nil {
			c.JSON(404, gin.H{"error": "Food not found"})
			return
		}
		_, err = models.GetOrderByID(orderItem.OrderID)
		if err != nil {
			c.JSON(404, gin.H{"error": "Order not found"})
			return
		}
		orderItem.CreatedAt = time.Now()
		orderItem.UpdatedAt = time.Now()
		err = models.CreateOrderItem(&orderItem)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Order item created successfully", "orderItem": orderItem})

	}

}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("orderItem_id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Invalid order item id"})
			return
		}
		var updatedOrderItem models.OrderItem
		if err := c.ShouldBindJSON(&updatedOrderItem); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		orderItem, err := models.GetOrderItem(idInt)
		if err != nil {
			c.JSON(404, gin.H{"error": "Order item not found"})
			return
		}
		if updatedOrderItem.Quantity != "" {
			orderItem.Quantity = updatedOrderItem.Quantity
		}

		if updatedOrderItem.UnitPrice != 0 {
			orderItem.UnitPrice = updatedOrderItem.UnitPrice
		}

		if updatedOrderItem.FoodID != 0 {
			//check if that Food with that ID exists
			_, err := models.GetFoodByID(updatedOrderItem.FoodID)
			if err != nil {
				c.JSON(404, gin.H{"error": "Food not found"})
				return
			}
			orderItem.FoodID = updatedOrderItem.FoodID

		}
		if updatedOrderItem.OrderID != 0 {
			//check if that Order with that ID exists
			_, err := models.GetOrderByID(updatedOrderItem.OrderID)
			if err != nil {
				c.JSON(404, gin.H{"error": "Order not found"})
				return
			}
			orderItem.OrderID = updatedOrderItem.OrderID
		}
		orderItem.UpdatedAt = time.Now()
		err = models.UpdateOrderItem(&orderItem)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Order item updated successfully", "orderItem": orderItem})

	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("order_id")
		orderId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Invalid order id"})
			return
		}
		orderItems, err := models.GetOrderItemsByOrder(orderId)
		if err != nil {
			c.JSON(404, gin.H{"error": "Order items not found"})
			return
		}
		c.JSON(200, gin.H{"orderItems": orderItems})

	}
}

// func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {

// }
