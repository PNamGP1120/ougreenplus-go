package article

import "time"

type Type string
type Status string

const (
	TypeArticle Type = "article"
	TypeGreen   Type = "greennews"
	TypeBlog    Type = "blog"

	StatusDraft  Status = "draft"
	StatusReview Status = "pending_review"
	StatusPub    Status = "published"
)

type Article struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `json:"title"`
	Summary     string     `json:"summary"`
	Content     string     `json:"content"`
	Thumbnail   string     `json:"thumbnail_url"`
	CategoryID  uint       `json:"category_id"`
	Type        Type       `gorm:"size:20" json:"type"`
	Status      Status     `gorm:"size:20" json:"status"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
