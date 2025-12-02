package blog

type CreateUpdateBlogDTO struct {
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
	Thumbnail string `json:"thumbnail_url"`
}
