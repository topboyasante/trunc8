package services

import (
	"context"
	"errors"

	"github.com/topboyasante/trunc8/internal/models"
	"github.com/topboyasante/trunc8/internal/repositories"
	"github.com/topboyasante/trunc8/internal/utils"
)

// This defines a new struct type called ShortnerService (like creating a blueprint).
// It doesn't initialize an instance - just defines what the struct looks like.
type ShortnerService struct {
	repository *repositories.ShortenerRepository // This field holds a pointer to a ShortenerRepository
}

// This function creates a new ShortnerService instance and returns a pointer to it.
// The *ShortnerService return type means it returns a pointer, not the struct value itself.
func NewShortnerService(repository *repositories.ShortenerRepository) *ShortnerService {
	// &ShortnerService{...} creates a new struct instance and returns its memory address (a pointer to it)
	return &ShortnerService{
		repository: repository,
	}
}

func (s *ShortnerService) ShortenURL(ctx context.Context, originalURL string) (*models.URL, error) {

	if originalURL == "" {
		return nil, errors.New("original URL cannot be empty")
	}

	encodedURL := utils.GenerateURLCode()

	// This creates a NEW instance of models.URL and returns a pointer to it.
	// It's not pointing to some existing URL struct in the models package - we're creating a fresh one here.
	url := &models.URL{
		OriginalURL: originalURL,
		Code:        encodedURL,
		ClickCount:  0,
	}

	// *url dereferences the pointer, converting it from *models.URL to models.URL
	// The Create function expects a models.URL value, not a pointer, so we use * to get the actual struct
	_, err := s.repository.Create(ctx, *url)

	if err != nil {
		return nil, err
	}

	return url, nil
}
