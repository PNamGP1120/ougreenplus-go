package media

import "time"

type Media struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	FileName   string    `json:"file_name"`
	FileSize   int64     `json:"file_size"`
	FileType   string    `json:"file_type"`
	URL        string    `json:"url"`
	UploadedBy uint      `json:"uploaded_by"`
	CreatedAt  time.Time `json:"created_at"`
}
