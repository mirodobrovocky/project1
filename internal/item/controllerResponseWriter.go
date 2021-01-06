package item

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeResponseOk(action string, responseWriter http.ResponseWriter, data interface{})  {
	writeResponse(action, responseWriter, http.StatusOK, data)
}

func writeResponse(action string, responseWriter http.ResponseWriter, successStatusCode int, data interface{})  {
	responseWriter.WriteHeader(successStatusCode)

	if err := json.NewEncoder(responseWriter).Encode(data); err != nil {
		log.Printf("action=%s status=internalSrverError error=%v", action, err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
