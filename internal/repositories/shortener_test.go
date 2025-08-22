package repositories

import (
	"context"
	"testing"

	"github.com/topboyasante/trunc8/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNewShortnerRepository(t *testing.T) {
	// Skip this test since it requires actual database connection
	// which is not available in testing environment
	t.Skip("Skipping NewShortnerRepository test - requires database connection")
}

func TestShortenerRepository_Create(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		repo := &ShortenerRepository{collection: mt.Coll}
		
		// Mock successful insert
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		
		url := models.URL{
			OriginalURL: "https://example.com",
			Code:        "TEST",
			ClickCount:  0,
		}
		
		id, err := repo.Create(context.Background(), url)
		
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		
		if id == "" {
			t.Error("Expected non-empty ID")
		}
	})
	
	mt.Run("insert error", func(mt *mtest.T) {
		repo := &ShortenerRepository{collection: mt.Coll}
		
		// Mock insert error
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   0,
			Code:    11000,
			Message: "duplicate key error",
		}))
		
		url := models.URL{
			OriginalURL: "https://example.com",
			Code:        "TEST",
			ClickCount:  0,
		}
		
		id, err := repo.Create(context.Background(), url)
		
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		
		if id != "" {
			t.Error("Expected empty ID when error occurs")
		}
	})
}

func TestShortenerRepository_FindOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		repo := &ShortenerRepository{collection: mt.Coll}
		
		// Mock successful find
		first := mtest.CreateCursorResponse(1, "trunc8-db.links", mtest.FirstBatch, bson.D{
			{"_id", "507f1f77bcf86cd799439011"},
			{"original_url", "https://example.com"},
			{"code", "TEST"},
			{"click_count", 5},
		})
		killCursor := mtest.CreateCursorResponse(0, "trunc8-db.links", mtest.NextBatch)
		mt.AddMockResponses(first, killCursor)
		
		url, err := repo.FindOne(context.Background(), "TEST")
		
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		
		if url == nil {
			t.Fatal("Expected non-nil URL")
		}
		
		if url.OriginalURL != "https://example.com" {
			t.Errorf("Expected original URL 'https://example.com', got '%s'", url.OriginalURL)
		}
		
		if url.Code != "TEST" {
			t.Errorf("Expected code 'TEST', got '%s'", url.Code)
		}
		
		if url.ClickCount != 5 {
			t.Errorf("Expected click count 5, got %d", url.ClickCount)
		}
	})
	
	mt.Run("not found", func(mt *mtest.T) {
		repo := &ShortenerRepository{collection: mt.Coll}
		
		// Mock no documents found
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "trunc8-db.links", mtest.FirstBatch))
		
		url, err := repo.FindOne(context.Background(), "NOTFOUND")
		
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		
		if url != nil {
			t.Error("Expected nil URL when not found")
		}
	})
	
	mt.Run("database error", func(mt *mtest.T) {
		repo := &ShortenerRepository{collection: mt.Coll}
		
		// Mock database error
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    1,
			Message: "database connection error",
		}))
		
		url, err := repo.FindOne(context.Background(), "TEST")
		
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		
		if url != nil {
			t.Error("Expected nil URL when error occurs")
		}
	})
}