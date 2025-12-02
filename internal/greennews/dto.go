package greennews

type CreateUpdateGreennewsDTO struct {
	Number string `json:"number"`
	Month  int    `json:"month"`
	Year   int    `json:"year"`
}
