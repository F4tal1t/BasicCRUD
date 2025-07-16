package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Security(c *fiber.Ctx) error {
	c.Response().Header.Add("Content-Security-Policy", "default-src 'self'")
	c.Response().Header.Add("X-Frame-Options", "DENY")
	c.Response().Header.Add("X-XSS-Protection", "1; mode=block")

	return c.Next()
}
