package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Properties struct {
	Uri 		string
	Database 	string
	Collection 	string
}

type ConnectionProvider interface {
	GetConnection() *mongo.Collection
}

type connectionProvider struct {
	connection *mongo.Collection
}

func (c connectionProvider) GetConnection() *mongo.Collection {
	return c.connection
}

func NewConnectionProvider(ctx context.Context, properties Properties) ConnectionProvider {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(properties.Uri))
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(properties.Database).Collection(properties.Collection)
	return connectionProvider{collection}
}
