package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func MiddlewareAuth(storage *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := storage.Get(c)
		if err != nil {
			panic(err)
		}
		userEmail := ""
		if email, ok := session.Get("userName").(string); ok {
			userEmail = email
		}
		c.Locals("userName", userEmail)
		return c.Next()
	}
}
