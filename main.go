package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/api", func(c *fiber.Ctx) error {
		c.Status(200)
		return c.SendString("Hello World")
	})
	log.Fatal(app.Listen("localhost:8000"))
}
