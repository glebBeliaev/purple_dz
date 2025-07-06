package pages

import (
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
	vacancies, err := h.repository.GetAll()
	if err != nil {
		h.cusstomLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(vacancies)
	return tadaptor.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
