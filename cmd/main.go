package main

import (
	"log"

	"github.com/KrishGarg/go-simple-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api.Post("/products", handlers.CreateProduct)
	api.Get("/products", handlers.GetAllProducts)

	log.Fatal(app.Listen(":3000"))
}
