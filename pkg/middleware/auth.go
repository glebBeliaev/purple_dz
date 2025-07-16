package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func MiddlewareAuth(storage *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := storage.Get(c)
		if err != nil {
			log.Printf("Ошибка получения сессии: %v", err)
			return c.Redirect("/login", fiber.StatusSeeOther)
		}
		userEmail := ""
		if email, ok := session.Get("userName").(string); ok {
			userEmail = email
		}
		c.Locals("userName", userEmail)
		return c.Next()
	}
}
