package database

import (
	"context"

	"github.com/topboyasante/trunc8/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBClient *mongo.Client

func ConnectToMongo(cfg *config.Config) error {
	ctx := context.TODO()
	dbClientOptions := options.Client().ApplyURI(cfg.Database.Url)

	client, err := mongo.Connect(ctx, dbClientOptions)

	if err != nil {
		return err
	}

	DBClient = client
	return nil
}

func DisconnectMongo() error {
	if DBClient != nil {
		return DBClient.Disconnect(context.TODO())
	}
	return nil
}
