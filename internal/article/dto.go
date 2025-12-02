package article

type CreateUpdateArticleDTO struct {
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
	Thumbnail  string `json:"thumbnail_url"`
	CategoryID uint   `json:"category_id"`
	Type       Type   `json:"type"`
	Status     Status `json:"status"`
}
