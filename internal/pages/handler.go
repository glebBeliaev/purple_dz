package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router        fiber.Router
	cusstomLogger *zerolog.Logger
}

func NewHomeHandler(router fiber.Router, customLogger *zerolog.Logger) *HomeHandler {
	return &HomeHandler{
		router:        router,
		cusstomLogger: customLogger,
	}

}
func (h *HomeHandler) Register() {
	h.router.Get("/", func(c *fiber.Ctx) error {
		h.cusstomLogger.Info().
			Bool("test", true).
			Str("test", "test").
			Int("test", 1).
			Msg("test")
		return c.SendString("Hello, World!")
	})
}
