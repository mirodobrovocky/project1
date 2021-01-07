package item

import (
	"context"
	"errors"
	"github.com/mirodobrovocky/project1/pkg/database"
	"github.com/mirodobrovocky/project1/pkg/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	FindAll() ([]Item, error)
	FindByName(name string) (*Item, error)
	Save(item Item) (*Item, error)
}

type repository struct {
	connectionProvider database.ConnectionProvider
}

func (r repository) FindAll() ([]Item, error) {
	ctx := context.TODO()

	cursor, err := r.connectionProvider.GetConnection().Find(ctx, bson.D{})
	if err != nil {
		return []Item{}, nil
	}

	defer cursor.Close(ctx)

	var items []Item
	if err = cursor.All(ctx, &items); err != nil {
		return []Item{}, nil
	}

	return items, nil
}

func (r repository) FindByName(name string) (*Item, error) {
	one, err := r.getOne("name", name)
	if err != nil {
		return nil, err
	}

	return one, nil
}

func (r repository) Save(item Item) (*Item, error) {
	ctx := context.TODO()

	insertOneResult, err := r.connectionProvider.GetConnection().InsertOne(ctx, item)
	if err != nil {
		return nil, resolveWriteError(err)
	}

	return r.getById(insertOneResult.InsertedID)
}

func (r repository) getById(id interface{}) (*Item, error) {
	return r.getOne("_id", id)
}

func (r repository) getOne(by string, search interface{}) (*Item, error) {
	ctx := context.TODO()

	result := r.connectionProvider.GetConnection().FindOne(ctx, bson.M{by: search})

	var item Item
	if err := result.Decode(&item); err != nil {
		return nil, resolveReadError(err)
	}

	return &item, nil
}

func NewRepository(connectionProvider database.ConnectionProvider) Repository {
	return repository{connectionProvider}
}

func resolveReadError(err error) error {
	if err == mongo.ErrNoDocuments {
		return exception.EntityNotFound
	}
	return err
}

func resolveWriteError(err error) error {
	if isDuplicate(err) {
		return exception.Conflict
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
