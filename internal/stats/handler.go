package stats

import (
	"github.com/gofiber/fiber/v2"

	"github.com/PNamGP1120/ougreenplus-go/internal/article"
	"github.com/PNamGP1120/ougreenplus-go/internal/database"
	"github.com/PNamGP1120/ougreenplus-go/internal/event"
	"github.com/PNamGP1120/ougreenplus-go/internal/newsletter"
)

type Stats struct {
	Articles    int64 `json:"articles"`
	Events      int64 `json:"events"`
	Subscribers int64 `json:"subscribers"`
}

func GetStats(c *fiber.Ctx) error {
	var s Stats

	database.DB.Model(&article.Article{}).Count(&s.Articles)
	database.DB.Model(&event.Event{}).Count(&s.Events)
	database.DB.Model(&newsletter.Subscriber{}).Count(&s.Subscribers)

	return c.JSON(s)
}
