package main

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func main() {
	app := fiber.New()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	app.Use(fiberzerolog.New())
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})
	app.Listen(":3000")
}
