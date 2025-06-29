package main

import (
	"github.com/glebbeliaev/purple_dz/config"
	pages "github.com/glebbeliaev/purple_dz/internal/pages/home"
	"github.com/glebbeliaev/purple_dz/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	logger := pkg.NewLogger(config.NewLogConfig())

	app := fiber.New()

	app.Use(recover.New())
	app.Use(slogfiber.New(logger))
	app.Static("/public", "./public")
	pages.NewHandler(app)

	logger.Info("Server started")
	app.Listen(":3000")
}
