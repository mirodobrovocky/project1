package item

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mirodobrovocky/project1/pkg/exception"
	"github.com/mirodobrovocky/project1/pkg/util"
	"net/http"
)

type CreateDto struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"gt=0"`
}

type ReadDto struct {
	Name  string  `json:"name"`
	Owner string  `json:"owner"`
	Price float64 `json:"price"`
}

type Controller interface {
	GetItems(response http.ResponseWriter, request *http.Request)
	GetItem(response http.ResponseWriter, request *http.Request)
	CreateItem(response http.ResponseWriter, request *http.Request)
}

type controller struct {
	service  Service
	validate *validator.Validate
}

func (c controller) GetItems(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	items, err := c.service.FindAll()
	if err != nil {
		handleInternalServerError("GetItems", response, err)
		return
	}

	var result []ReadDto
	for _, item := range items {
		result = append(result, ReadDto{Name: item.Name, Owner: item.Owner, Price: item.Price})
	}
	writeResponseOk("GetItems", response, result)
}

func (c controller) GetItem(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	name := params["name"]
	if item, err := c.service.FindByName(name); err == nil {
		writeResponseOk("GetItem", response, ReadDto{
			Name: item.Name,
			Owner: item.Owner,
			Price: item.Price,
		})
	} else if err == exception.EntityNotFound {
		handleNotFound("GetItem", response, err)
	} else {
		handleInternalServerError("GetItem", response, err)
	}
}

func (c controller) CreateItem(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var item CreateDto

	if err:= util.JsonDecode(request.Body, &item); err != nil {
		handleBadRequest("CreateItem", response, err)
		return
	}

	if err := c.validate.Struct(item); err != nil {
		handleValidationErrors("CreateItem", response, err.(validator.ValidationErrors))
		return
	}

	save, err := c.service.Create(item); if err != nil {
		if err == exception.Conflict {
			handleConflict("CreateItem", response, err)
		} else {
			handleInternalServerError("CreateItem", response, err)
		}
		return
	}

	writeResponse("CreateItem", response, http.StatusCreated, ReadDto{
		Name:  save.Name,
		Owner: save.Owner,
		Price: save.Price,
	})
}

func NewController(service Service, validate *validator.Validate) Controller {
	return controller{service: service, validate: validate}
}
