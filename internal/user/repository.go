package user

import (
	"github.com/mirodobrovocky/project1/pkg/database"
)

type Repository interface {
	FindById(id int) (*User, error)
}

type repository struct {
	connection database.Connection
}

func (r repository) FindById(id int) (*User, error) {
	var user User
	if err:= r.connection.FindById(&user, id); err != nil {
		return nil, err
	}
	return &user, nil
}

func NewRepository(connection database.Connection) Repository {
	return repository{connection: connection}
}
