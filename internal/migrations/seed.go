package migrations

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
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

// SeedAll seeds ~10 sample data for each table
func SeedAll(db *gorm.DB) {
	SeedAdmin(db)
	SeedUsers(db)
	SeedCategories(db)
	SeedTags(db)
	SeedArticles(db)
	SeedBlogs(db)
	SeedGreennews(db)
	SeedEvents(db)
	SeedMedia(db)
	SeedSubscribers(db)
}

/*
	============================================
	  USERS

============================================
*/
func SeedAdmin(db *gorm.DB) {
	var count int64
	db.Model(&user.User{}).Count(&count)

	if count == 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), 12)

		admin := user.User{
			Email:    "admin@ou.edu.vn",
			Password: string(hash),
			Role:     user.RoleAdmin,
		}

		db.Create(&admin)
		log.Println("🌱 Seeded: default admin created")
	}
}

func SeedUsers(db *gorm.DB) {
	var count int64
	db.Model(&user.User{}).Where("email != ?", "admin@ou.edu.vn").Count(&count)

	if count > 0 {
		return
	}

	for i := 1; i <= 10; i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte("123456"), 12)
		db.Create(&user.User{
			Email:    fmt.Sprintf("user%d@ou.edu.vn", i),
			Password: string(hash),
			Role:     user.RoleEditor,
		})
	}

	log.Println("🌱 Seeded: users (10 records)")
}

/*
	============================================
	  CATEGORY

============================================
*/
func SeedCategories(db *gorm.DB) {
	var count int64
	db.Model(&category.Category{}).Count(&count)
	if count > 0 {
		return
	}

	names := []string{"Môi trường", "Công nghệ xanh", "Giáo dục", "Năng lượng", "Hành động", "Biển & Đại dương", "Khí hậu", "Tái chế", "Sống xanh"}
	for _, name := range names {
		db.Create(&category.Category{
			Name:        name,
			Description: "Chuyên mục về " + name,
		})
	}

	log.Println("🌱 Seeded: categories (10 records)")
}

/*
	============================================
	  TAGS

============================================
*/
func SeedTags(db *gorm.DB) {
	var count int64
	db.Model(&tag.Tag{}).Count(&count)
	if count > 0 {
		return
	}

	names := []string{"green", "environment", "ocean", "energy", "climate", "campus", "university", "student", "event"}
	for _, name := range names {
		db.Create(&tag.Tag{Name: name})
	}

	log.Println("🌱 Seeded: tags (10 records)")
}

/*
	============================================
	  ARTICLES

============================================
*/
func SeedArticles(db *gorm.DB) {
	var count int64
	db.Model(&article.Article{}).Count(&count)
	if count > 0 {
		return
	}

	for i := 1; i <= 10; i++ {
		db.Create(&article.Article{
			Title:      fmt.Sprintf("Bài viết mẫu %d", i),
			Summary:    "Lorem ipsum dolor sit amet",
			Content:    "<p>Nội dung bài viết lorem ipsum...</p>",
			CategoryID: uint((i % 8) + 1),
			Type:       article.TypeArticle,
			Status:     article.StatusPub,
			PublishedAt: func() *time.Time {
				t := time.Now().AddDate(0, 0, -i)
				return &t
			}(),
		})
	}

	log.Println("🌱 Seeded: articles (10 records)")
}

/*
	============================================
	  BLOG

============================================
*/
func SeedBlogs(db *gorm.DB) {
	var count int64
	db.Model(&blog.Blog{}).Count(&count)
	if count > 0 {
		return
	}

	for i := 1; i <= 10; i++ {
		db.Create(&blog.Blog{
			Title:     fmt.Sprintf("Câu chuyện xanh %d", i),
			Summary:   "Blog summary lorem ipsum",
			Content:   "<p>Câu chuyện xanh...</p>",
			Thumbnail: "",
		})
	}

	log.Println("🌱 Seeded: blogs (10 records)")
}

/*
	============================================
	  GREENNEWS

============================================
*/
func SeedGreennews(db *gorm.DB) {
	var count int64
	db.Model(&greennews.Greennews{}).Count(&count)
	if count > 0 {
		return
	}

	for i := 1; i <= 10; i++ {
		db.Create(&greennews.Greennews{
			Number: fmt.Sprintf("0%d/2025", i),
			Month:  i,
			Year:   2025,
		})
	}

	log.Println("🌱 Seeded: greennews (10 records)")
}

/*
	============================================
	  EVENTS

============================================
*/
func SeedEvents(db *gorm.DB) {
	var count int64
	db.Model(&event.Event{}).Count(&count)
	if count > 0 {
		return
	}

	for i := 1; i <= 10; i++ {
		db.Create(&event.Event{
			Title:       fmt.Sprintf("Sự kiện %d", i),
			Description: "Mô tả sự kiện",
			StartDate:   time.Now().AddDate(0, 0, i),
			EndDate:     time.Now().AddDate(0, 0, i+1),
			Location:    "OU Campus",
			Status:      event.EventUpcoming,
		})
	}

	log.Println("🌱 Seeded: events (10 records)")
}

/*
	============================================
	  MEDIA

============================================
*/
func SeedMedia(db *gorm.DB) {
	var count int64
	db.Model(&media.Media{}).Count(&count)
	if count > 0 {
		return
	}

	for i := 1; i <= 10; i++ {
		db.Create(&media.Media{
			FileName:   fmt.Sprintf("image%d.jpg", i),
			FileSize:   12345,
			FileType:   "image/jpeg",
			URL:        fmt.Sprintf("https://example.com/media/image%d.jpg", i),
			UploadedBy: 1,
		})
	}

	log.Println("🌱 Seeded: media (10 records)")
}

/*
	============================================
	  NEWSLETTER SUBSCRIBERS

============================================
*/
func SeedSubscribers(db *gorm.DB) {
	var count int64
	db.Model(&newsletter.Subscriber{}).Count(&count)
	if count > 0 {
		return
	}

	for i := 1; i <= 10; i++ {
		db.Create(&newsletter.Subscriber{
			Email:    fmt.Sprintf("subscriber%d@gmail.com", i),
			IsActive: true,
		})
	}

	log.Println("🌱 Seeded: subscribers (10 records)")
}
