package db

import (
	"FirstAPI/internal/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDBService struct {
	Client *mongo.Client
}

func NewMongoDBService() (*MongoDBService, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DATABASE_URI))

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return &MongoDBService{Client: client}, nil
}

func (s *MongoDBService) GetCollection(database, collection string) *mongo.Collection {
	return s.Client.Database(database).Collection(collection)
}
