package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/MichaelTinney/zapmanejo-cleanbackend/internal/database"
	"github.com/MichaelTinney/zapmanejo-cleanbackend/internal/middleware"
	"github.com/MichaelTinney/zapmanejo-cleanbackend/internal/models"
)

func SetupAnimalRoutes(app *fiber.App) {
	animal := app.Group("/api/animals", middleware.JWTProtected())

	animal.Get("/", getAnimals)
	animal.Post("/", createAnimal)
	animal.Put("/:id", updateAnimal)
	animal.Delete("/:id", deleteAnimal)
}

func getAnimals(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	var animals []models.Animal
	database.DB.Where("user_id = ?", userID).Find(&animals)
	return c.JSON(animals)
}

func createAnimal(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	var animal models.Animal
	if err := c.BodyParser(&animal); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	animal.UserID = userID
	database.DB.Create(&animal)
	return c.JSON(animal)
}

func updateAnimal(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	id := c.Params("id")
	var animal models.Animal
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&animal).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Animal not found"})
	}
	c.BodyParser(&animal)
	database.DB.Save(&animal)
	return c.JSON(animal)
}

func deleteAnimal(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	id := c.Params("id")
	database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Animal{})
	return c.Status(200).JSON(fiber.Map{"success": true})
}
