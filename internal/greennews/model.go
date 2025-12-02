package greennews

import "time"

type Greennews struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Number    string    `json:"number"`
	Month     int       `json:"month"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
