package main

import (
	"fmt"
	"musictracks/db"
	"musictracks/routes"
	"os"

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

	port := os.Getenv("MUSICTRACK_APP_PORT")

	// Set a default value if the environment variable is not set
	if port == "" {
		fmt.Printf("Error: Environment variable %s is not set.\n", "MUSICTRACK_APP_PORT")
		os.Exit(1)
	}

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	if clientID == "" {
		fmt.Printf("Error: Environment variable %s is not set.\n", "SPOTIFY_CLIENT_ID")
		os.Exit(1)
	}

	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if clientSecret == "" {
		fmt.Printf("Error: Environment variable %s is not set.\n", "SPOTIFY_CLIENT_SECRET")
		os.Exit(1)
	}

	r.Run(":" + port)
}
