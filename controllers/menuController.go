package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/models"

	"github.com/gin-gonic/gin"
)

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {

		menus, err := models.GetMenus()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"menus": menus})

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("menu_id")

		menuId, err := strconv.Atoi(id) // Convert to int here
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
			return
		}

		menu, err := models.GetMenuByID(menuId)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{"menu": menu})

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		var menu models.Menu
           
		if err := c.ShouldBindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//check if end date is after the start date
		if menu.StartDate.After(menu.EndDate) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Start date cannot be after end date"})
            return
        }



		menu.CreatedAt = time.Now()
		menu.UpdatedAt = time.Now()

		fmt.Printf("After setting times:\n%+v\n", menu)

		err := models.CreateMenu(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Menu created successfully",
			"menu":    menu,
		})

	}
}

func UpdateMenu() gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("menu_id")
        idInt, err := strconv.Atoi(id)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
            return
        }

        var updatedMenu models.Menu
        if err := c.ShouldBindJSON(&updatedMenu); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        menu, err := models.GetMenuByID(idInt)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
            return
        }

        if updatedMenu.Name != "" {
            menu.Name = updatedMenu.Name
        }

        if updatedMenu.Category != "" {
            menu.Category = updatedMenu.Category
        }

		if !updatedMenu.StartDate.IsZero() {
			menu.StartDate = updatedMenu.StartDate
		}
		
		if !updatedMenu.EndDate.IsZero() {
			menu.EndDate = updatedMenu.EndDate
		}

		//check if end date is after the start date
		if menu.StartDate.After(menu.EndDate) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Start date cannot be after end date"})
            return
        }
		

        menu.UpdatedAt = time.Now()

        err = models.UpdateMenu(*menu)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Menu updated successfully",
            "menu":    menu,
        })
    }
}


func DeleteMenu() gin.HandlerFunc {
    return func(c *gin.Context) {

        id := c.Param("menu_id")

        menuId, err := strconv.Atoi(id) // Convert to int here
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
            return
        }

        // Delete related food items
        err = models.DeleteFoodItemsByMenuID(menuId)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete related food items"})
            return
        }
        err = models.DeleteMenu(menuId)
        if err != nil {
            if err == sql.ErrNoRows {
                c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            }
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Menu and related food items deleted successfully"})

    }
}
