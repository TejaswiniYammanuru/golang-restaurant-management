package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/models"
	"github.com/gin-gonic/gin"
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		foods, err := models.GetAllFoods()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"foods": foods})
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("food_id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food ID"})
			return
		}
		food, err := models.GetFoodByID(idInt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"food": food})
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var food models.Food

		if err := c.ShouldBindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := models.GetMenuByID(food.MenuID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Menu not found"})
			return
		}

		food.CreatedAt = time.Now()
		food.UpdatedAt = time.Now()

		err = models.CreateFood(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Food created successfully",
			"food":    food,
		})
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("food_id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food ID"})
			return
		}

		var updatedFood models.Food
		if err := c.ShouldBindJSON(&updatedFood); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		food, err := models.GetFoodByID(idInt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
			return
		}

		if updatedFood.Name != "" {
			food.Name = updatedFood.Name
		}

		if updatedFood.MenuID != 0 {
			_, err := models.GetMenuByID(updatedFood.MenuID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Menu not found"})
				return
			}
			food.MenuID = updatedFood.MenuID
		}

		if updatedFood.Price != 0 {
			food.Price = updatedFood.Price
		}

		if updatedFood.FoodImage != "" {
			food.FoodImage = updatedFood.FoodImage
		}
		food.UpdatedAt = time.Now()

		err = models.UpdateFood(food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Food updated successfully",
			"food":    food,
		})
	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("food_id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food ID"})
			return
		}

		err = models.DeleteFood(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})
	}
}
