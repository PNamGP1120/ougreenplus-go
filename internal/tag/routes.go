package tag

import (
	"github.com/gofiber/fiber/v2"

	"github.com/PNamGP1120/ougreenplus-go/internal/config"
	"github.com/PNamGP1120/ougreenplus-go/internal/middleware"
)

func RegisterRoutes(r fiber.Router, cfg *config.Config) {
	h := NewHandler()

	r.Get("/", h.List)

	admin := r.Group("", middleware.JWTProtected(cfg))
	admin.Post("/", h.Create)
	admin.Put("/:id", h.Update)
	admin.Delete("/:id", h.Delete)
}
