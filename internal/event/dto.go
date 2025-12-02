package event

import "time"

type CreateUpdateEventDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Location    string    `json:"location"`
	PosterURL   string    `json:"poster_url"`
	Status      Status    `json:"status"`
}

type RegisterDTO struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	StudentID string `json:"student_id"`
}
