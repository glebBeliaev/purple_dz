package main

import (
	"log"

	"github.com/glebbeliaev/purple_dz/config"
	"github.com/glebbeliaev/purple_dz/internal/pages"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	dbconfig := config.NewDataBaseConfig()
	log.Println(dbconfig)
	app := fiber.New()
	app.Use(recover.New())

	homeHandler := pages.NewHomeHandler(app)
	homeHandler.Register()
	app.Listen(":3000")
}
