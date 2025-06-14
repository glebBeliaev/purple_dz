package pages

import (
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
	h := HomeHandler{
		router: router,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	users := []User{
		{Id: 1, Name: "Anton"},
		{Id: 2, Name: "Vasia"},
	}
	names := []string{"Anton", "Vasia"}
	data := struct {
		Names []string
		Users []User
	}{
		Names: names,
		Users: users,
	}
	return c.Render("page", data)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Error")
}
