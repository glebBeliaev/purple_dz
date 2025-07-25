package createnews

import (
	"github.com/glebbeliaev/purple_dz/internal/registration"
	"github.com/glebbeliaev/purple_dz/pkg/tadapter"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type CreateNewsHandler struct {
	router       fiber.Router
	store        *session.Store
	customLogger *zerolog.Logger
	repository   *registration.UserRepository
}

func NewHandler(router fiber.Router, store *session.Store, customLogger *zerolog.Logger, repository *registration.UserRepository) {
	h := &CreateNewsHandler{
		router:       router,
		store:        store,
		customLogger: customLogger,
		repository:   repository,
	}
	h.router.Get("/createNews", h.createNews)
}

func (h *CreateNewsHandler) createNews(c *fiber.Ctx) error {
	userName := c.Locals("userName").(string)
	if userName != "" {
		component := views.CreateNews()
		return tadapter.Render(c, component, fiber.StatusOK)
	}
	return c.Redirect("/login", fiber.StatusSeeOther)
}
