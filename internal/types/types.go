package types

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
}
