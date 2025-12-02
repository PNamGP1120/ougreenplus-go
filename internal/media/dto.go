package media

type UploadResponse struct {
	ID       uint   `json:"id"`
	URL      string `json:"url"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileSize int64  `json:"file_size"`
}
