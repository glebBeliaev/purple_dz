package main

import (
	"github.com/glebbeliaev/purple_dz/config"
	pages "github.com/glebbeliaev/purple_dz/internal/pages/home"
	"github.com/glebbeliaev/purple_dz/internal/pages/register"
	"github.com/glebbeliaev/purple_dz/internal/registration"
	"github.com/glebbeliaev/purple_dz/pkg/logger"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))

	app.Use(recover.New())
	app.Static("/public", "./public")

	pages.NewHandler(app)
	register.NewHandler(app)
	registration.NewHandler(app)

	customLogger.Info().Msg("Server started")

	app.Listen(":3000")
}
