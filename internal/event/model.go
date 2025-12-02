package event

import "time"

type Status string

const (
	EventUpcoming Status = "upcoming"
	EventOngoing  Status = "ongoing"
	EventFinished Status = "finished"
)

type Event struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Location    string    `json:"location"`
	PosterURL   string    `json:"poster_url"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Registration struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	EventID   uint      `json:"event_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	StudentID string    `json:"student_id"`
	CreatedAt time.Time `json:"created_at"`
}
