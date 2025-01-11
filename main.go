package main

import (
	log "log"
	os "os"

	fiber "github.com/gofiber/fiber/v2"
	config "github.com/herumitra/apt-auth/config"
	router "github.com/herumitra/apt-auth/router"
	godotenv "github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get port from environment
	serverPort := os.Getenv("SERVER_PORT")

	// Initialize database
	if err := config.SetupDB(); err != nil {
		log.Fatal(err)
	}

	// Initialize app
	app := fiber.New()

	// Setup routes
	router.SetupRoutes(app)

	// Start app
	log.Fatal(app.Listen(":" + serverPort))
}
