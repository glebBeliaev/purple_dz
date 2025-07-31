package newsAnimal

import (
	"github.com/glebbeliaev/purple_dz/internal/models"
	"github.com/glebbeliaev/purple_dz/internal/news"
	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type NewsAnimalHandler struct {
	router        fiber.Router
	storage       *session.Store
	repository    *news.NewsRepository
	cusstomLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, storage *session.Store, logger *zerolog.Logger, repository *news.NewsRepository) {
	h := &NewsAnimalHandler{
		router:        router,
		storage:       storage,
		cusstomLogger: logger,
		repository:    repository,
	}

	h.router.Get("/animal", h.newsAnimal)
}

func (h *NewsAnimalHandler) newsAnimal(c *fiber.Ctx) error {
	animalBlock, err := models.BuildCategoryBlock(c, h.repository, "Животные", "Животные", "animalPage", 8)
	if err != nil {
		h.cusstomLogger.Error().Msg("Ошибка при загрузке категории Животные: " + err.Error())
		return c.SendStatus(500)
	}

	component := views.NewsFood(animalBlock)
	return tadapter.Render(c, component, fiber.StatusOK)
}
