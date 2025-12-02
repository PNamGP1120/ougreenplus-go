package middleware

import "github.com/gofiber/fiber/v2"

func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).JSON(fiber.Map{
			"success": false,
			"error":   e.Message,
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}
