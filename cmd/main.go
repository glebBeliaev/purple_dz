package main

import (
	"github.com/glebbeliaev/purple_dz/config"
	"github.com/glebbeliaev/purple_dz/internal/pages"
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
	homeHandler := pages.NewHomeHandler(app)
	homeHandler.Register()
	logger.Info("Server started")
	app.Listen(":3000")
}
