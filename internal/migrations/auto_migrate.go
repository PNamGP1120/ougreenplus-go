package migrations

import (
	"gorm.io/gorm"

	"github.com/PNamGP1120/ougreenplus-go/internal/article"
	"github.com/PNamGP1120/ougreenplus-go/internal/blog"
	"github.com/PNamGP1120/ougreenplus-go/internal/category"
	"github.com/PNamGP1120/ougreenplus-go/internal/event"
	"github.com/PNamGP1120/ougreenplus-go/internal/greennews"
	"github.com/PNamGP1120/ougreenplus-go/internal/media"
	"github.com/PNamGP1120/ougreenplus-go/internal/newsletter"
	"github.com/PNamGP1120/ougreenplus-go/internal/tag"
	"github.com/PNamGP1120/ougreenplus-go/internal/user"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&blog.Blog{},
		&category.Category{},
		&tag.Tag{},
		&event.Event{},
		&event.Registration{},
		&greennews.Greennews{},
		&newsletter.Subscriber{},
		&media.Media{},
	)
}
