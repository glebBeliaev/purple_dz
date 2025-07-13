package register

import (
	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type RegisterHandler struct {
	router  fiber.Router
	storage *session.Store
}

func NewHandler(router fiber.Router, store *session.Store) {
	h := &RegisterHandler{
		router:  router,
		storage: store,
	}
	h.router.Get("/register", h.register)
	h.router.Get("/login", h.login)
}

func (h *RegisterHandler) register(c *fiber.Ctx) error {
	component := views.Register()
	return tadapter.Render(c, component, fiber.StatusOK)
}

func (h *RegisterHandler) login(c *fiber.Ctx) error {
	component := views.Login()
	return tadapter.Render(c, component, fiber.StatusOK)
}
