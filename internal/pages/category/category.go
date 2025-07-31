package category

import (
	"github.com/a-h/templ"
	"github.com/glebbeliaev/purple_dz/internal/models"
	"github.com/glebbeliaev/purple_dz/internal/news"
	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type PageCategory struct {
	Route      string
	Title      string
	Category   string
	QueryParam string
	ViewFunc   func(block models.CategoryBlock) templ.Component
}

func NewHandler(router fiber.Router, storage *session.Store, logger *zerolog.Logger, repository *news.NewsRepository) {
	categories := []PageCategory{
		{
			Route:      "/food",
			Title:      "Еда",
			Category:   "Еда",
			QueryParam: "foodPage",
			ViewFunc:   views.NewsCategoryPage,
		},
		{
			Route:      "/cars",
			Title:      "Машины",
			Category:   "Машины",
			QueryParam: "carsPage",
			ViewFunc:   views.NewsCategoryPage,
		},
		{
			Route:      "/animal",
			Title:      "Животные",
			Category:   "Животные",
			QueryParam: "animalPage",
			ViewFunc:   views.NewsCategoryPage,
		},
		{
			Route:      "/sport",
			Title:      "Спорт",
			Category:   "Спорт",
			QueryParam: "sportPage",
			ViewFunc:   views.NewsCategoryPage,
		},
		{
			Route:      "/music",
			Title:      "Музыка",
			Category:   "Музыка",
			QueryParam: "musicPage",
			ViewFunc:   views.NewsCategoryPage,
		},
		{
			Route:      "/music",
			Title:      "Музыка",
			Category:   "Музыка",
			QueryParam: "musicPage",
			ViewFunc:   views.NewsCategoryPage,
		},
		{
			Route:      "/tech",
			Title:      "Технологии",
			Category:   "Технологии",
			QueryParam: "techPage",
			ViewFunc:   views.NewsCategoryPage,
		},
		{
			Route:      "/other",
			Title:      "Прочее",
			Category:   "Прочее",
			QueryParam: "otherPage",
			ViewFunc:   views.NewsCategoryPage,
		},
	}

	for _, cat := range categories {
		router.Get(cat.Route, buildHandler(cat, repository, logger))
	}
}

func buildHandler(cat PageCategory, repo *news.NewsRepository, logger *zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		block, err := models.BuildCategoryBlock(c, repo, cat.Category, cat.Title, cat.QueryParam, 8)
		if err != nil {
			logger.Error().Msgf("Ошибка при загрузке категории %s: %s", cat.Category, err.Error())
			return c.SendStatus(500)
		}
		return tadapter.Render(c, cat.ViewFunc(block), 200)
	}
}
