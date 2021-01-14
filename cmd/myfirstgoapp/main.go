package main

import (
	_ "embed"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mirodobrovocky/project1/internal/config"
	"github.com/mirodobrovocky/project1/internal/item"
	"github.com/mirodobrovocky/project1/internal/user"
	"github.com/mirodobrovocky/project1/pkg/database"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
)

func main() {
	//go:embed "properties.yml"
	var propertiesBytes []byte
	var properties config.Properties
	if err := yaml.Unmarshal(propertiesBytes, &properties); err != nil {
		log.Fatal(err)
	}
	mongoConnection := database.GetMongoConnection(properties)
	postgresConnection := database.GetPostgresConnection(properties)

	userRepository := user.NewRepository(postgresConnection)
	userService := user.NewService(userRepository)
	itemsRepository := item.NewRepository(mongoConnection)
	itemsService := item.NewService(itemsRepository, userService)
	itemsController := item.NewController(itemsService, validator.New())

	r := mux.NewRouter()

	r.HandleFunc("/items", itemsController.GetItems).Methods("GET")
	r.HandleFunc("/items/{name:[a-zA-Z][a-zA-Z0-9_]*}", itemsController.GetItem).Methods("GET")
	r.HandleFunc("/items", itemsController.CreateItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}
