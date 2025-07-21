package createnews

import (
	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type CreateNewsHandler struct {
	router  fiber.Router
	storage *session.Store
}

func NewHandler(router fiber.Router, store *session.Store) {
	h := &CreateNewsHandler{
		router:  router,
		storage: store,
	}
	h.router.Get("/createNews", h.createNews)
}

func (h *CreateNewsHandler) createNews(c *fiber.Ctx) error {
	component := views.CreateNews()
	return tadapter.Render(c, component, fiber.StatusOK)
}
