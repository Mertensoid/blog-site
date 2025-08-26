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
		if email, ok := session.Get("email").(string); ok {
			userEmail = email
		}
		userName := ""
		if name, ok := session.Get("name").(string); ok {
			userName = name
		}
		c.Locals("email", userEmail)
		c.Locals("name", userName)
		return c.Next()
	}
}
