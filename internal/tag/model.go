package tag

type Tag struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex" json:"name"`
	Slug string `json:"slug"`
}
