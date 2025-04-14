package main

import (
	"log"

	"github.com/DAT-CANDIDATE/db"
	"github.com/DAT-CANDIDATE/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// MongoDB connection
	db.ConnectDB("mongodb+srv://verifyhire:vh12345@verifyhire.outzf6f.mongodb.net/?retryWrites=true&w=majority&appName=verifyhire")

	// Initialize Gin
	router := gin.Default()

	// Setup Routes
	routes.SetupRoutes(router)

	// Start server
	log.Println("Server is running on port 8081")
	log.Fatal(router.Run(":8081"))

	router.Run(":8081")
}
