package pages

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

type Catalog struct {
	CategoryName string
}

func NewHomeHandler(router fiber.Router) *HomeHandler {
	h := &HomeHandler{router: router}
	h.registerRoutes()
	return h
}

func (h *HomeHandler) registerRoutes() {
	h.router.Get("/", h.handleHome)
}

func (h *HomeHandler) handleHome(c *fiber.Ctx) error {
	categories := []string{"#Еда", "#Животные", "#Спорт", "#Музыка", "#Технологии", "#Другое"}
	return c.Render("index", fiber.Map{
		"Categories": categories,
	})
}
