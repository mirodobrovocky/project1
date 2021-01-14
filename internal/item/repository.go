package item

import (
	"github.com/mirodobrovocky/project1/pkg/database"
)

type Repository interface {
	FindAll() ([]Item, error)
	FindByName(name string) (*Item, error)
	Save(item Item) (*Item, error)
}

type repository struct {
	db database.Connection
}

func (r repository) FindAll() ([]Item, error) {
	var items []Item
	if err:= r.db.FindAll(&items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r repository) FindByName(name string) (*Item, error) {
	var item Item
	if err:= r.db.FindOne(&item, database.Filter{Field: "name", Value: name}); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r repository) Save(item Item) (*Item, error) {
	var saved Item
	if err := r.db.Save(&saved, &item); err != nil {
		return nil, err
	}
	return &saved, nil
}

func NewRepository(db database.Connection) Repository {
	return repository{db}
}
