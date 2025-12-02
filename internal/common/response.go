package common

import "github.com/gofiber/fiber/v2"

func Success(data interface{}) fiber.Map {
	return fiber.Map{
		"success": true,
		"data":    data,
	}
}

func Error(message string) fiber.Map {
	return fiber.Map{
		"success": false,
		"error":   message,
	}
}

func Pagination(data interface{}, page, size int, total int64) fiber.Map {
	return fiber.Map{
		"success": true,
		"data":    data,
		"page":    page,
		"size":    size,
		"total":   total,
	}
}
