package services

import (
	"encoding/base64"
	"errors"
	"strconv"
	"time"

	"github.com/topboyasante/trunc8/internal/models"
)

func ShortenURL(originalURL string) (*models.URL, error) {

	if originalURL == "" {
		return nil, errors.New("original URL cannot be empty")
	}

	id := int(time.Now().UnixMilli())
	idStr := strconv.Itoa(id)
	encodedURL := base64.StdEncoding.EncodeToString([]byte(idStr))

	url := &models.URL{
		ID:          id,
		OriginalURL: originalURL,
		Code:        encodedURL,
		ClickCount:  0,
		CreatedAt:   time.Now(),
	}

	return url, nil
}
