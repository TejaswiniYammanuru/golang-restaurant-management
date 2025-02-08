package controllers

import (
	"strconv"
	"time"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/models"
	"github.com/gin-gonic/gin"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		tables, err := models.GetTables()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"tables": tables})

	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("table_id")
		tableId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Invalid table id"})
			return
		}
		table, err := models.GetTableByID(tableId)
		if err != nil {
			c.JSON(404, gin.H{"error": "Table not found"})
			return
		}
		c.JSON(200, gin.H{"table": table})

	}

}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table
		if err := c.ShouldBindJSON(&table); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		table.CreatedAt = time.Now()
		table.UpdatedAt = time.Now()
		err := models.CreateTable(&table)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Table created successfully", "table": table})

	}

}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("table_id")
		tableId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Invalid table id"})
			return
		}

		var updatedTable models.Table
		if err := c.ShouldBindJSON(&updatedTable); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		table, err := models.GetTableByID(tableId)
		if err != nil {
			c.JSON(404, gin.H{"error": "Table not found"})
			return
		}
		if updatedTable.NumberOfGuests != 0 {
			table.NumberOfGuests = updatedTable.NumberOfGuests

		}
		if updatedTable.TableNumber != 0 {
			table.TableNumber = updatedTable.TableNumber

		}

		table.UpdatedAt = time.Now()
		err = models.UpdateTable(table)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Table updated successfully", "table": table})

	}
}
