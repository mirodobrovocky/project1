package main

import (
	"context"
	_ "embed"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mirodobrovocky/project1/internal/item"
	"github.com/mirodobrovocky/project1/pkg/database"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
)

type properties struct {
	Database struct {
		Uri			string
		Database 	string
		Collection 	string
	}
}

func main() {
	//go:embed "properties.yml"
	var propertiesBytes []byte
	var properties properties
	err := yaml.Unmarshal(propertiesBytes, &properties); if err != nil {
		log.Fatal(err)
	}
	databaseConnectionProvider := database.NewConnectionProvider(context.TODO(), database.Properties{
		Uri:        properties.Database.Uri,
		Database:   properties.Database.Database,
		Collection: properties.Database.Collection,
	})

	itemsRepository := item.NewRepository(databaseConnectionProvider)
	itemsService := item.NewService(itemsRepository)
	itemsController := item.NewController(itemsService, validator.New())

	r := mux.NewRouter()

	r.HandleFunc("/items", itemsController.GetItems).Methods("GET")
	r.HandleFunc("/items/{name:[a-zA-Z][a-zA-Z0-9_]*}", itemsController.GetItem).Methods("GET")
	r.HandleFunc("/items", itemsController.CreateItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}
