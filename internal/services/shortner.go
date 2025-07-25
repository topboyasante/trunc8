package services

import (
	"errors"
	"time"

	"github.com/topboyasante/trunc8/internal/models"
	"github.com/topboyasante/trunc8/internal/utils"
)

func ShortenURL(originalURL string) (*models.URL, error) {

	if originalURL == "" {
		return nil, errors.New("original URL cannot be empty")
	}

	id := int(time.Now().UnixMilli())
	encodedURL := utils.GenerateURLCode()

	url := &models.URL{
		ID:          id,
		OriginalURL: originalURL,
		Code:        encodedURL,
		ClickCount:  0,
		CreatedAt:   time.Now(),
	}

	return url, nil
}
