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

	}

}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("food_id")
		idInt, err := strconv.Atoi(id)
		if err!= nil {
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

		menuID, err := strconv.Atoi(food.MenuID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid MenuID"})
			return
		}

		_, err = models.GetMenuByID(menuID)
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

// func round(num float64) int {

// }

// func toFixed(num float64, precision int) float64 {
// }

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}

}
