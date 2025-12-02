package media

import (
	"github.com/gofiber/fiber/v2"

	"github.com/PNamGP1120/ougreenplus-go/internal/config"
	"github.com/PNamGP1120/ougreenplus-go/internal/middleware"
)

func RegisterRoutes(r fiber.Router, cfg *config.Config) {
	h := NewHandler(cfg)

	admin := r.Group("", middleware.JWTProtected(cfg))

	admin.Post("/upload", h.Upload)
	admin.Get("/", h.List)
	admin.Get("/:id", h.Get)
	admin.Delete("/:id", h.Delete)
}
