package newsletter

type SubscribeDTO struct {
	Email string `json:"email"`
}

type SendMailDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
