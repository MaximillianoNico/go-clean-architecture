package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) PublicRoute(c *fiber.Ctx) error {
	return c.Next()
}
