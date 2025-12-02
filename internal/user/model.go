package user

import "time"

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleEditor Role = "editor"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Password  string    `json:"-"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
