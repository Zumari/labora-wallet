package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseJson(response http.ResponseWriter, status int, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("error while marshalling object %v, trace: %+v", data, err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	_, err = response.Write(bytes)
	if err != nil {
		return fmt.Errorf("error while writing bytes to response writer: %+v", err)
	}

	return nil
}
