package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/PNamGP1120/ougreenplus-go/internal/article"
	"github.com/PNamGP1120/ougreenplus-go/internal/auth"
	"github.com/PNamGP1120/ougreenplus-go/internal/blog"
	"github.com/PNamGP1120/ougreenplus-go/internal/category"
	"github.com/PNamGP1120/ougreenplus-go/internal/config"
	"github.com/PNamGP1120/ougreenplus-go/internal/database"
	"github.com/PNamGP1120/ougreenplus-go/internal/event"
	"github.com/PNamGP1120/ougreenplus-go/internal/greennews"
	"github.com/PNamGP1120/ougreenplus-go/internal/media"
	"github.com/PNamGP1120/ougreenplus-go/internal/middleware"
	"github.com/PNamGP1120/ougreenplus-go/internal/migrations"
	"github.com/PNamGP1120/ougreenplus-go/internal/newsletter"
	"github.com/PNamGP1120/ougreenplus-go/internal/stats"
	"github.com/PNamGP1120/ougreenplus-go/internal/tag"
	"github.com/PNamGP1120/ougreenplus-go/internal/user"
)

func main() {
	// Load configs
	cfg := config.Load()

	// Connect DB
	database.Connect(cfg)

	// Run migrations (NO import cycles)
	migrations.AutoMigrate(database.DB)
	migrations.SeedAll(database.DB)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// Global middlewares
	app.Use(middleware.Logger)
	app.Use(middleware.CORS)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Register all routes
	auth.RegisterRoutes(app.Group("/auth"), cfg)
	user.RegisterRoutes(app.Group("/users"), cfg)
	article.RegisterRoutes(app.Group("/articles"), cfg)
	category.RegisterRoutes(app.Group("/categories"), cfg)
	tag.RegisterRoutes(app.Group("/tags"), cfg)
	blog.RegisterRoutes(app.Group("/blog"), cfg)
	event.RegisterRoutes(app.Group("/events"), cfg)
	greennews.RegisterRoutes(app.Group("/greennews"), cfg)
	newsletter.RegisterRoutes(app.Group("/newsletter"), cfg)
	media.RegisterRoutes(app.Group("/media"), cfg)
	stats.RegisterRoutes(app.Group("/admin"), cfg)

	log.Printf("🚀 Server running on port %s", cfg.Port)

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
