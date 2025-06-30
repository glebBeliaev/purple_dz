package register

import (
	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &RegisterHandler{
		router: router,
	}
	h.router.Get("/register", h.register)
}

func (h *RegisterHandler) register(c *fiber.Ctx) error {
	component := views.Register()
	return tadapter.Render(c, component)
}
