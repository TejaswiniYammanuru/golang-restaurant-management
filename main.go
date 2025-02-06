package main

import (
	"log"
	"os"

	"github.com/TejaswiniYammanuru/golang-restaurant-management/database"
	"github.com/TejaswiniYammanuru/golang-restaurant-management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Initialize the database once
	database.InitDB()

	router := gin.New()
	router.Use(gin.Logger())

	// Register routes
	// routes.UserRoutes(router)
	// router.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	// routes.TableRoutes(router)
	// routes.OrderRoutes(router)
	// routes.OrderItemRoutes(router)
	// routes.InvoiceRoutes(router)

	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
