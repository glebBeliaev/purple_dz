package pages

import (
	"math"
	"net/http"

	"github.com/glebbeliaev/purple_dz/internal/vacancy"
	"github.com/glebbeliaev/purple_dz/pkg/tadaptor"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router        fiber.Router
	cusstomLogger *zerolog.Logger
	repository    *vacancy.VacancyRepository
}

type User struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLoger *zerolog.Logger, repository *vacancy.VacancyRepository) {
	h := &HomeHandler{
		router:        router,
		cusstomLogger: customLoger,
		repository:    repository,
	}

	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)
	count := int(math.Ceil(float64(h.repository.CountAll() / PAGE_ITEMS)))
	vacancies, err := h.repository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.cusstomLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(vacancies, count, page)
	return tadaptor.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
