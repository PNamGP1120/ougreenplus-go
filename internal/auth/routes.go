package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/PNamGP1120/ougreenplus-go/internal/config"
	"github.com/PNamGP1120/ougreenplus-go/internal/middleware"
)

func RegisterRoutes(r fiber.Router, cfg *config.Config) {
	h := NewHandler(cfg)

	r.Post("/login", h.Login)
	r.Post("/refresh", h.Refresh)
	r.Post("/logout", h.Logout)

	protected := r.Group("", middleware.JWTProtected(cfg))
	protected.Get("/me", h.Me)
}
