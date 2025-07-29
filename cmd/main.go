package main

import (
	"time"

	"github.com/glebbeliaev/purple_dz/config"
	"github.com/glebbeliaev/purple_dz/internal/news"
	createnews "github.com/glebbeliaev/purple_dz/internal/pages/createNews"
	pages "github.com/glebbeliaev/purple_dz/internal/pages/home"
	"github.com/glebbeliaev/purple_dz/internal/pages/register"
	"github.com/glebbeliaev/purple_dz/internal/registration"
	"github.com/glebbeliaev/purple_dz/pkg/database"
	"github.com/glebbeliaev/purple_dz/pkg/logger"
	"github.com/glebbeliaev/purple_dz/pkg/middleware"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
)

func main() {
	config.Init()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	dbConfig := config.NewDataBaseConfig()

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))

	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()

	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.MiddlewareAuth(store))

	app.Use(recover.New())
	app.Static("/public", "./public")

	userRepo := registration.NewUserRepository(dbpool)
	newsRepo := news.NewNewsRepository(dbpool)

	//Handlers
	pages.NewHandler(app, store, customLogger, newsRepo)
	register.NewHandler(app, store)
	registration.NewHandler(app, customLogger, userRepo, store)
	createnews.NewHandler(app, store, customLogger, userRepo)
	news.NewHandler(app, customLogger, newsRepo, store)
	customLogger.Info().Msg("Server started")

	app.Listen(":3000")
}
