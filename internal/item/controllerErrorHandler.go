package item

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func handleNotFound(action string, responseWriter http.ResponseWriter, err error) {
	responseWriter.WriteHeader(http.StatusNotFound)
	log.Printf("action=%s status=notFound error=%v", action, err)
}

func handleConflict(action string, responseWriter http.ResponseWriter, err error) {
	responseWriter.WriteHeader(http.StatusConflict)
	log.Printf("action=%s status=conflict error=%v", action, err)
}

func handleBadRequest(action string, responseWriter http.ResponseWriter, err error) {
	responseWriter.WriteHeader(http.StatusBadRequest)
	log.Printf("action=%s status=badRequest error=%v", action, err)
}

func handleInternalServerError(action string, responseWriter http.ResponseWriter, err error) {
	responseWriter.WriteHeader(http.StatusInternalServerError)
	log.Printf("action=%s status=internalServerError error=%v", action, err)
}

func handleValidationErrors(action string, responseWriter http.ResponseWriter, validationErrors validator.ValidationErrors) {
	responseWriter.WriteHeader(http.StatusBadRequest)
	log.Printf("action=%s status=badRequest error=validationErrors", action)
	for _, err := range validationErrors {
		errorMsg := fmt.Sprintf("error=validationErrors field=%s validation=%s value=%v", err.Field(), err.ActualTag(), err.Value())
		_, _ = responseWriter.Write([]byte(errorMsg))
		log.Printf(errorMsg)
	}
}
