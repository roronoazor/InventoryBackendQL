package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/roronoazor/InventoryBackendQL/database"
)

func main() {
	log.Println("Loading Server")
	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8000"))
}
