package newsFood

import (
	"math"

	"github.com/glebbeliaev/purple_dz/internal/models"
	"github.com/glebbeliaev/purple_dz/internal/news"
	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type NewsFoodHandler struct {
	router        fiber.Router
	storage       *session.Store
	repository    *news.NewsRepository
	cusstomLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, storage *session.Store, logger *zerolog.Logger, repository *news.NewsRepository) {
	h := &NewsFoodHandler{
		router:        router,
		storage:       storage,
		cusstomLogger: logger,
		repository:    repository,
	}

	h.router.Get("/food", h.newsFood)
}

func (h *NewsFoodHandler) newsFood(c *fiber.Ctx) error {
	PAGE_ITEMS := 8

	// --- Популярные (по категории)
	popularPage := c.QueryInt("popularPage", 1)

	popularTotal := h.repository.CountByCategory("Популярное")
	popularCount := int(math.Ceil(float64(popularTotal) / float64(PAGE_ITEMS)))

	popularNews, err := h.repository.GetByCategory("Популярное", PAGE_ITEMS, (popularPage-1)*PAGE_ITEMS)
	if err != nil {
		h.cusstomLogger.Error().Msg("Ошибка при загрузке популярных новостей: " + err.Error())
		return c.SendStatus(500)
	}

	foodBlock := models.CategoryBlock{
		Title:      "Популярное",
		News:       popularNews,
		Page:       popularPage,
		Count:      popularCount,
		QueryParam: "popularPage",
	}
	component := views.NewsFood(foodBlock)
	return tadapter.Render(c, component, fiber.StatusOK)
}
