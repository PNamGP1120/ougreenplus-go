package blog

import "time"

type Blog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	Content   string    `json:"content"`
	Thumbnail string    `json:"thumbnail_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
