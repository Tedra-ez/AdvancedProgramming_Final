package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	client   *mongo.Client
	database string
}

const defaultDatabase = "clothes_store"

func NewMongoDBClient(ctx context.Context, uri string) (*MongoDBClient, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	log.Println("MongoDB connected")
	return &MongoDBClient{client: client, database: defaultDatabase}, nil
}

func (c *MongoDBClient) Database() *mongo.Database {
	return c.client.Database(c.database)
}

func (c *MongoDBClient) Collection(name string) *mongo.Collection {
	return c.Database().Collection(name)
}

func (c *MongoDBClient) Close(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}
