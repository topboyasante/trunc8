package models

// struct field names should be exported (start with uppercase)
// so they can be accessed outside the package

type URL struct {
	ID          string `bson:"_id,omitempty"`
	OriginalURL string `bson:"original_url"`
	Code        string `bson:"code"`
	ClickCount  int    `bson:"click_count"`
}
