package main

import (
	"time"

	"github.com/glebbeliaev/purple_dz/config"
	"github.com/glebbeliaev/purple_dz/internal/pages"
	"github.com/glebbeliaev/purple_dz/internal/sitemap"
	"github.com/glebbeliaev/purple_dz/internal/vacancy"
	"github.com/glebbeliaev/purple_dz/middleware"
	"github.com/glebbeliaev/purple_dz/pkg/database"
	"github.com/glebbeliaev/purple_dz/pkg/logger"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
)

func main() {
	config.Init()
	config.NewDataBaseConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDataBaseConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")
	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()

	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{Storage: storage})
	app.Use(middleware.AuthMiddleware(store))

	//REPOSITORIES
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, customLogger)

	//HANDLERS
	pages.NewHandler(app, customLogger, vacancyRepo, store)
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	sitemap.NewHandler(app)

	app.Listen(":3000")
}
