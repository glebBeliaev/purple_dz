package pages

import (
	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	router fiber.Router
}

type Catalog struct {
	CategoryName string
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}

	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main("World")
	return tadapter.Render(c, component)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
