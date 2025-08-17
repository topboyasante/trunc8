package repositories

import (
	"context"

	"github.com/topboyasante/trunc8/internal/database"
	"github.com/topboyasante/trunc8/internal/models"
	"go.mongodb.org/mongo-driver/bson"
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
	return result.InsertedID.(string), nil
}

func (r *ShortenerRepository) FindOne(ctx context.Context, id string) (*models.URL, error) {
	var user models.URL

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
