package repositories

import (
	"context"
	"fmt"

	"github.com/topboyasante/trunc8/internal/database"
	"github.com/topboyasante/trunc8/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShortenerRepository struct {
	collection *mongo.Collection
}

func NewShortnerRepository() *ShortenerRepository {
	collection := database.DBClient.Database("trunc8-db").Collection("links")
	return &ShortenerRepository{
		collection: collection,
	}
}

func (r *ShortenerRepository) Create(ctx context.Context, user models.URL) (string, error) {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	// Convert primitive.ObjectID to string using Hex()
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("expected ObjectID for InsertedID, got %T", result.InsertedID)
	}
	return id.Hex(), nil
}

func (r *ShortenerRepository) FindOne(ctx context.Context, code string) (*models.URL, error) {
	var user models.URL

	err := r.collection.FindOne(ctx, bson.M{"code": code}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
