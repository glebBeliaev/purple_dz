package pages

import (
	"fmt"

	"github.com/glebbeliaev/purple_dz/pkg/tadaptor"
	"github.com/glebbeliaev/purple_dz/views"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router        fiber.Router
	cusstomLogger *zerolog.Logger
}

type User struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router) {
	fmt.Println(">>> routes registered")
	h := &HomeHandler{
		router: router,
	}

	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	fmt.Println(">>> home called")
	component := views.Main()
	return tadaptor.Render(c, component)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
