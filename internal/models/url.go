package models

import "time"

// struct field names should be exported (start with uppercase)
// so they can be accessed outside the package

type URL struct {
	ID          int
	OriginalURL string
	Code        string
	ClickCount  int
	CreatedAt   time.Time
}
