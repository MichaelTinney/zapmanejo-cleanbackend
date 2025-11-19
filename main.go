package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/MichaelTinney/zapmanejo-cleanbackend/internal/database"
	"github.com/MichaelTinney/zapmanejo-cleanbackend/internal/routes"
)

func main() {
	// Load .env file (optional, production uses env vars directly)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to PostgreSQL
	database.Connect()

	app := fiber.New(fiber.Config{
		AppName: "ZapManejo v1.0",
	})

	// CORS middleware - configurable via environment variable
	allowOrigins := os.Getenv("CORS_ALLOW_ORIGINS")
	if allowOrigins == "" {
		allowOrigins = "*"
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Health endpoint for cloud platforms
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": "healthy",
			"service": "zapmanejo-backend",
		})
	})

	// Setup all application routes
	routes.Setup(app)

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Start server in a goroutine for graceful shutdown
	go func() {
		log.Printf("ZapManejo backend starting on port %s", port)
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Shutting down server gracefully...")
	
	// Get shutdown timeout from environment or default to 30 seconds
	shutdownTimeout := 30 * time.Second
	if timeoutStr := os.Getenv("SHUTDOWN_TIMEOUT_SECONDS"); timeoutStr != "" {
		if timeout, err := strconv.Atoi(timeoutStr); err == nil {
			shutdownTimeout = time.Duration(timeout) * time.Second
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited successfully")
}
