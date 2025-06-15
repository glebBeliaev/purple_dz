package main

import (
	"github.com/glebbeliaev/purple_dz/config"
	"github.com/glebbeliaev/purple_dz/internal/pages"
	"github.com/glebbeliaev/purple_dz/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	logger := pkg.NewLogger(config.NewLogConfig())

	engine := html.New("./html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(recover.New())
	app.Use(slogfiber.New(logger))
	homeGroup := app.Group("/")
	pages.NewHomeHandler(homeGroup)

	logger.Info("Server started")
	app.Listen(":3000")
}
