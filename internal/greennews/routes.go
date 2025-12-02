package greennews

import (
	"github.com/gofiber/fiber/v2"

	"github.com/PNamGP1120/ougreenplus-go/internal/config"
	"github.com/PNamGP1120/ougreenplus-go/internal/middleware"
)

func RegisterRoutes(r fiber.Router, cfg *config.Config) {
	h := NewHandler()

	// public
	r.Get("/", h.List)
	r.Get("/:id", h.Get)

	// admin
	admin := r.Group("", middleware.JWTProtected(cfg))
	admin.Post("/", h.Create)
	admin.Put("/:id", h.Update)
	admin.Delete("/:id", h.Delete)
}
