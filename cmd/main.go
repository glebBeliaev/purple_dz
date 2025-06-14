package main

import (
	"strings"

	"github.com/glebbeliaev/purple_dz/config"
	"github.com/glebbeliaev/purple_dz/internal/pages"
	"github.com/glebbeliaev/purple_dz/pkg/logger"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	config.Init()
	config.NewDataBaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	engine := html.New("./html", ".html")
	engine.AddFuncMap(map[string]interface{}{
		"ToUpper": func(s string) string {
			return strings.ToUpper(s)
		},
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())

	pages.NewHandler(app)

	app.Listen(":3000")
}
