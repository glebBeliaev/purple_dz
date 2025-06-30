package registration

import (
	"github.com/a-h/templ"

	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/views/components"
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := RegisterHandler{
		router: router,
	}
	api := h.router.Group("/api")
	api.Post("/signUp", h.signUp)
}

func (h *RegisterHandler) signUp(c *fiber.Ctx) error {
	email := c.FormValue("email")
	var component templ.Component
	if email == "" {
		component = components.Notification("Не задан email", components.NotificationFail)
		return tadapter.Render(c, component)
	}
	component = components.Notification("Vacancy created!", components.NotificationSuccess)
	return tadapter.Render(c, component)
}
