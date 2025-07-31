package pages

import (
	"github.com/glebbeliaev/purple_dz/internal/models"
	"github.com/glebbeliaev/purple_dz/internal/news"
	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router        fiber.Router
	storage       *session.Store
	repository    *news.NewsRepository
	cusstomLogger *zerolog.Logger
}

type Catalog struct {
	CategoryName string
}

func NewHandler(router fiber.Router, storage *session.Store, logger *zerolog.Logger, repository *news.NewsRepository) {
	h := &HomeHandler{
		router:        router,
		storage:       storage,
		cusstomLogger: logger,
		repository:    repository,
	}

	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	latestBlock, err := models.BuildCategoryBlock(c, h.repository, "", "Последние новости", "page", 4)
	if err != nil {
		h.cusstomLogger.Error().Msg("Ошибка при загрузке последних новостей: " + err.Error())
		return c.SendStatus(500)
	}

	popularBlock, err := models.BuildCategoryBlock(c, h.repository, "Популярное", "Популярное", "popularPage", 4)
	if err != nil {
		h.cusstomLogger.Error().Msg("Ошибка при загрузке популярных новостей: " + err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(latestBlock, popularBlock)
	return tadapter.Render(c, component, fiber.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
