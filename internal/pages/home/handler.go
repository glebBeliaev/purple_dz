package pages

import (
	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type HomeHandler struct {
	router  fiber.Router
	storage *session.Store
}

type Catalog struct {
	CategoryName string
}

func NewHandler(router fiber.Router, storage *session.Store) {
	h := &HomeHandler{
		router:  router,
		storage: storage,
	}

	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return tadapter.Render(c, component, fiber.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
