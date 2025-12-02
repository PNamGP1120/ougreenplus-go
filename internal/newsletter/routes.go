package newsletter

import (
	"github.com/gofiber/fiber/v2"

	"github.com/PNamGP1120/ougreenplus-go/internal/config"
	"github.com/PNamGP1120/ougreenplus-go/internal/middleware"
)

func RegisterRoutes(r fiber.Router, cfg *config.Config) {
	h := NewHandler()

	r.Post("/subscribe", h.Subscribe)
	r.Post("/unsubscribe", h.Unsubscribe)

	admin := r.Group("", middleware.JWTProtected(cfg))
	admin.Get("/subscribers", h.List)
	admin.Post("/send", h.Send)
}
