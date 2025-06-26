package vacancy

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router        fiber.Router
	cusstomLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, cusstomLogger *zerolog.Logger) {
	h := &VacancyHandler{
		router:        router,
		cusstomLogger: cusstomLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	return c.SendString("Vacancy created!")
}
