package pages

import (
	"math"

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
	PAGE_ITEMS := 4
	page := c.QueryInt("page", 1)

	count := int(math.Ceil(float64(h.repository.CountAll() / PAGE_ITEMS)))
	news, err := h.repository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.cusstomLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(news, count, page)
	return tadapter.Render(c, component, fiber.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
