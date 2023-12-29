package main

import (
	"musictracks/db"
	"musictracks/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files (index.html)
	r.StaticFile("/", "./index.html")

	// Initialize database connection
	db.InitDB()

	// Setup routes
	routes.SetupRoutes(r)

	r.Run(":8000")
}
