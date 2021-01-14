package database

import (
	"github.com/mirodobrovocky/project1/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

type postgresConnection struct {
	connection *gorm.DB
}

func (p postgresConnection) FindById(dest interface{}, id interface{}) error {
	p.connection.First(dest, id)
	return p.connection.Error
}

func (p postgresConnection) FindOne(dest interface{}, filter Filter) error {
	panic("implement me")
}

func (p postgresConnection) FindAll(dest interface{}) error {
	panic("implement me")
}

func (p postgresConnection) Save(dest interface{}, save interface{}) error {
	panic("implement me")
}

func GetPostgresConnection(properties config.Properties) Connection {
	db, err := gorm.Open(postgres.Open(properties.Postgres.Dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: properties.Postgres.SingularTable,
		},
	})

	if err != nil {
		log.Panicf("no db error=%v", err)
	}

	return postgresConnection{db}
}

