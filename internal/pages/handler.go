package pages

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

func NewHomeHandler(router fiber.Router) *HomeHandler {
	return &HomeHandler{
		router: router,
	}

}
func (h *HomeHandler) Register() {
	h.router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
