package services

import (
	"context"
	"errors"
	"testing"

	"github.com/topboyasante/trunc8/internal/models"
)

// Mock repository for testing
type mockShortenerRepository struct {
	createFunc  func(ctx context.Context, url models.URL) (string, error)
	findOneFunc func(ctx context.Context, code string) (*models.URL, error)
}

func (m *mockShortenerRepository) Create(ctx context.Context, url models.URL) (string, error) {
	if m.createFunc != nil {
		return m.createFunc(ctx, url)
	}
	return "mock-id", nil
}

func (m *mockShortenerRepository) FindOne(ctx context.Context, code string) (*models.URL, error) {
	if m.findOneFunc != nil {
		return m.findOneFunc(ctx, code)
	}
	return &models.URL{
		ID:          "mock-id",
		OriginalURL: "https://example.com",
		Code:        code,
		ClickCount:  0,
	}, nil
}

func TestNewShortnerService(t *testing.T) {
	mockRepo := &mockShortenerRepository{}
	service := NewShortnerService(mockRepo)
	
	if service == nil {
		t.Error("Expected service to be non-nil")
	}
	
	if service.repository == nil {
		t.Error("Expected service repository to be non-nil")
	}
}

func TestShortenURL_Success(t *testing.T) {
	mockRepo := &mockShortenerRepository{
		createFunc: func(ctx context.Context, url models.URL) (string, error) {
			if url.OriginalURL != "https://example.com" {
				t.Errorf("Expected original URL to be 'https://example.com', got '%s'", url.OriginalURL)
			}
			if len(url.Code) != 4 {
				t.Errorf("Expected code length to be 4, got %d", len(url.Code))
			}
			if url.ClickCount != 0 {
				t.Errorf("Expected click count to be 0, got %d", url.ClickCount)
			}
			return "test-id", nil
		},
	}
	
	service := NewShortnerService(mockRepo)
	ctx := context.Background()
	
	result, err := service.ShortenURL(ctx, "https://example.com")
	
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	
	if result == nil {
		t.Fatal("Expected result to be non-nil")
	}
	
	if result.ID != "test-id" {
		t.Errorf("Expected ID to be 'test-id', got '%s'", result.ID)
	}
	
	if result.OriginalURL != "https://example.com" {
		t.Errorf("Expected original URL to be 'https://example.com', got '%s'", result.OriginalURL)
	}
	
	if len(result.Code) != 4 {
		t.Errorf("Expected code length to be 4, got %d", len(result.Code))
	}
	
	if result.ClickCount != 0 {
		t.Errorf("Expected click count to be 0, got %d", result.ClickCount)
	}
}

func TestShortenURL_EmptyURL(t *testing.T) {
	mockRepo := &mockShortenerRepository{}
	service := NewShortnerService(mockRepo)
	ctx := context.Background()
	
	result, err := service.ShortenURL(ctx, "")
	
	if err == nil {
		t.Fatal("Expected error for empty URL")
	}
	
	if result != nil {
		t.Error("Expected result to be nil when error occurs")
	}
	
	expectedError := "original URL cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}

func TestShortenURL_RepositoryError(t *testing.T) {
	mockRepo := &mockShortenerRepository{
		createFunc: func(ctx context.Context, url models.URL) (string, error) {
			return "", errors.New("database error")
		},
	}
	
	service := NewShortnerService(mockRepo)
	ctx := context.Background()
	
	result, err := service.ShortenURL(ctx, "https://example.com")
	
	if err == nil {
		t.Fatal("Expected error from repository")
	}
	
	if result != nil {
		t.Error("Expected result to be nil when error occurs")
	}
	
	if err.Error() != "database error" {
		t.Errorf("Expected error 'database error', got '%s'", err.Error())
	}
}

func TestRedirectURL_Success(t *testing.T) {
	mockRepo := &mockShortenerRepository{
		findOneFunc: func(ctx context.Context, code string) (*models.URL, error) {
			if code != "TEST" {
				t.Errorf("Expected code to be 'TEST', got '%s'", code)
			}
			return &models.URL{
				ID:          "test-id",
				OriginalURL: "https://example.com",
				Code:        "TEST",
				ClickCount:  5,
			}, nil
		},
	}
	
	service := NewShortnerService(mockRepo)
	ctx := context.Background()
	
	result, err := service.RedirectURL(ctx, "TEST")
	
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	
	if result != "https://example.com" {
		t.Errorf("Expected result to be 'https://example.com', got '%s'", result)
	}
}

func TestRedirectURL_EmptyCode(t *testing.T) {
	mockRepo := &mockShortenerRepository{}
	service := NewShortnerService(mockRepo)
	ctx := context.Background()
	
	result, err := service.RedirectURL(ctx, "")
	
	if err == nil {
		t.Fatal("Expected error for empty code")
	}
	
	if result != "" {
		t.Error("Expected result to be empty string when error occurs")
	}
	
	expectedError := "code cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}

func TestRedirectURL_RepositoryError(t *testing.T) {
	mockRepo := &mockShortenerRepository{
		findOneFunc: func(ctx context.Context, code string) (*models.URL, error) {
			return nil, errors.New("not found")
		},
	}
	
	service := NewShortnerService(mockRepo)
	ctx := context.Background()
	
	result, err := service.RedirectURL(ctx, "NOTFOUND")
	
	if err == nil {
		t.Fatal("Expected error from repository")
	}
	
	if result != "" {
		t.Error("Expected result to be empty string when error occurs")
	}
	
	if err.Error() != "not found" {
		t.Errorf("Expected error 'not found', got '%s'", err.Error())
	}
}