package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AuthMiddleware(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := store.Get(c)
		if err != nil {
			panic(err)
		}
		userEmail := ""
		if name, ok := session.Get("email").(string); ok {
			userEmail = name
		}
		c.Locals("email", userEmail)
		return c.Next()
	}
}
