package database

import (
	"context"
	"errors"
	"github.com/mirodobrovocky/project1/internal/config"
	"github.com/mirodobrovocky/project1/pkg/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type mongoConnection struct {
	connection *mongo.Collection
}

func (m mongoConnection) FindById(dest interface{}, id interface{}) error {
	return m.getOne(dest, Filter{Field: "_id", Value: id})
}

func (m mongoConnection) FindOne(dest interface{}, filter Filter) error {
	return m.getOne(dest, filter)
}

func (m mongoConnection) FindAll(dest interface{}) error {
	ctx := context.Background()
	cursor, err := m.connection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, dest); err != nil {
		return err
	}

	return nil
}

func (m mongoConnection) Save(dest interface{}, save interface{}) error {
	ctx := context.Background()

	if beforeSaveAction, ok := save.(BeforeSaveAction); ok {
		beforeSaveAction.BeforeSave()
	}
	insertOneResult, err := m.connection.InsertOne(ctx, save)
	if err != nil {
		return resolveWriteError(err)
	}

	return m.FindById(dest, insertOneResult.InsertedID)
}

func (m mongoConnection) getOne(dest interface{}, filter Filter) error {
	ctx := context.Background()

	result := m.connection.FindOne(ctx, bson.M{filter.Field: filter.Value})

	if err := result.Decode(dest); err != nil {
		return resolveReadError(err)
	}

	return nil
}

func GetMongoConnection(properties config.Properties) Connection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(properties.Mongo.Uri))
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(properties.Mongo.Database).Collection(properties.Mongo.Collection)
	return mongoConnection{collection}
}

func resolveWriteError(err error) error {
	if isDuplicate(err) {
		return exception.Conflict
	}

	return err
}

func resolveReadError(err error) error {
	if err == mongo.ErrNoDocuments {
		return exception.EntityNotFound
	}
	return err
}

func isDuplicate(err error) bool {
	var writeException mongo.WriteException
	if errors.As(err, &writeException) {
		for _, we := range writeException.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}

	return false
}
