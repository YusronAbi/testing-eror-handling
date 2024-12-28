package main

import (
	"inventory-management/database"
	"inventory-management/routes"
)

func main() {
	// Initialize Database
	database.Connect()

	// Setup Routes
	router := routes.SetupRoutes()

	// Run Server
	router.Run(":8080")
}
