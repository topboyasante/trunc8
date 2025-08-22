package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/topboyasante/trunc8/internal/models"
	"github.com/topboyasante/trunc8/internal/types"
)

// Mock service for testing
type mockShortnerService struct {
	shortenURLFunc  func(ctx context.Context, originalURL string) (*models.URL, error)
	redirectURLFunc func(ctx context.Context, code string) (string, error)
}

func (m *mockShortnerService) ShortenURL(ctx context.Context, originalURL string) (*models.URL, error) {
	if m.shortenURLFunc != nil {
		return m.shortenURLFunc(ctx, originalURL)
	}
	return &models.URL{
		ID:          "test-id",
		OriginalURL: originalURL,
		Code:        "TEST",
		ClickCount:  0,
	}, nil
}

func (m *mockShortnerService) RedirectURL(ctx context.Context, code string) (string, error) {
	if m.redirectURLFunc != nil {
		return m.redirectURLFunc(ctx, code)
	}
	return "https://example.com", nil
}

func TestNewShortnerHandler(t *testing.T) {
	mockService := &mockShortnerService{}
	handler := NewShortnerHandler(mockService)
	
	if handler == nil {
		t.Error("Expected handler to be non-nil")
	}
	
	if handler.service == nil {
		t.Error("Expected handler service to be non-nil")
	}
}

func TestShortenURL_Success(t *testing.T) {
	mockService := &mockShortnerService{
		shortenURLFunc: func(ctx context.Context, originalURL string) (*models.URL, error) {
			if originalURL != "https://example.com" {
				t.Errorf("Expected URL 'https://example.com', got '%s'", originalURL)
			}
			return &models.URL{
				ID:          "test-id",
				OriginalURL: originalURL,
				Code:        "TEST",
				ClickCount:  0,
			}, nil
		},
	}
	
	handler := NewShortnerHandler(mockService)
	
	payload := types.ShortenRequest{URL: "https://example.com"}
	jsonPayload, _ := json.Marshal(payload)
	
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	
	handler.ShortenURL(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
	
	var response models.URL
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error unmarshaling response: %v", err)
	}
	
	if response.OriginalURL != "https://example.com" {
		t.Errorf("Expected original URL 'https://example.com', got '%s'", response.OriginalURL)
	}
	
	if response.Code != "TEST" {
		t.Errorf("Expected code 'TEST', got '%s'", response.Code)
	}
}

func TestShortenURL_InvalidJSON(t *testing.T) {
	handler := NewShortnerHandler(&mockShortnerService{})
	
	req := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	
	handler.ShortenURL(w, req)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
	
	if !strings.Contains(w.Body.String(), "Invalid JSON") {
		t.Error("Expected 'Invalid JSON' in response body")
	}
}

func TestShortenURL_ServiceError(t *testing.T) {
	mockService := &mockShortnerService{
		shortenURLFunc: func(ctx context.Context, originalURL string) (*models.URL, error) {
			return nil, errors.New("service error")
		},
	}
	
	handler := NewShortnerHandler(mockService)
	
	payload := types.ShortenRequest{URL: "https://example.com"}
	jsonPayload, _ := json.Marshal(payload)
	
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	
	handler.ShortenURL(w, req)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
	
	if !strings.Contains(w.Body.String(), "Error shortening url") {
		t.Error("Expected 'Error shortening url' in response body")
	}
}

func TestShortenURL_UnsupportedMethod(t *testing.T) {
	handler := NewShortnerHandler(&mockShortnerService{})
	
	req := httptest.NewRequest(http.MethodGet, "/shorten", nil)
	w := httptest.NewRecorder()
	
	handler.ShortenURL(w, req)
	
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}
	
	if !strings.Contains(w.Body.String(), "Not found") {
		t.Error("Expected 'Not found' in response body")
	}
}

func TestRedirectURL_Success(t *testing.T) {
	mockService := &mockShortnerService{
		redirectURLFunc: func(ctx context.Context, code string) (string, error) {
			if code != "TEST" {
				t.Errorf("Expected code 'TEST', got '%s'", code)
			}
			return "https://example.com", nil
		},
	}
	
	handler := NewShortnerHandler(mockService)
	
	req := httptest.NewRequest(http.MethodGet, "/TEST", nil)
	w := httptest.NewRecorder()
	
	handler.RedirectURL(w, req)
	
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, got %d", http.StatusMovedPermanently, w.Code)
	}
	
	location := w.Header().Get("Location")
	if location != "https://example.com" {
		t.Errorf("Expected location 'https://example.com', got '%s'", location)
	}
}

func TestRedirectURL_NoShortCode(t *testing.T) {
	handler := NewShortnerHandler(&mockShortnerService{})
	
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	
	handler.RedirectURL(w, req)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
	
	if !strings.Contains(w.Body.String(), "No short code provided") {
		t.Error("Expected 'No short code provided' in response body")
	}
}

func TestRedirectURL_ServiceError(t *testing.T) {
	mockService := &mockShortnerService{
		redirectURLFunc: func(ctx context.Context, code string) (string, error) {
			return "", errors.New("not found")
		},
	}
	
	handler := NewShortnerHandler(mockService)
	
	req := httptest.NewRequest(http.MethodGet, "/NOTFOUND", nil)
	w := httptest.NewRecorder()
	
	handler.RedirectURL(w, req)
	
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
	
	if !strings.Contains(w.Body.String(), "Unable to retrieve full url") {
		t.Error("Expected 'Unable to retrieve full url' in response body")
	}
}

func TestRedirectURL_UnsupportedMethod(t *testing.T) {
	handler := NewShortnerHandler(&mockShortnerService{})
	
	req := httptest.NewRequest(http.MethodPost, "/TEST", nil)
	w := httptest.NewRecorder()
	
	handler.RedirectURL(w, req)
	
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}
	
	if !strings.Contains(w.Body.String(), "Not found") {
		t.Error("Expected 'Not found' in response body")
	}
}