package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/PNamGP1120/ougreenplus-go/internal/config"
	"github.com/PNamGP1120/ougreenplus-go/internal/middleware"
)

func RegisterRoutes(r fiber.Router, cfg *config.Config) {
	admin := r.Group("", middleware.JWTProtected(cfg))

	admin.Get("/", List)
	admin.Get("/:id", Get)
	admin.Post("/", Create)
	admin.Put("/:id", Update)
	admin.Delete("/:id", Delete)
}
