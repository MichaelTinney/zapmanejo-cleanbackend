package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/MichaelTinney/zapmanejo-cleanbackend/internal/database"
	"github.com/MichaelTinney/zapmanejo-cleanbackend/internal/routes"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	// Connect to PostgreSQL
	database.Connect()

	app := fiber.New(fiber.Config{
		AppName: "ZapManejo v1.0",
	})

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Routes
	routes.Setup(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Printf("ZapManejo backend running on :%s", port)
	log.Fatal(app.Listen(":" + port))
}
