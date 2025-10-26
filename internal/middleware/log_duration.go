package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogDuration() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)
		println("Request duration:", duration.String())
		return err
	}
}
