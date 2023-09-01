package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type RestAPIError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NO_ERROR() RestAPIError {
	return RestAPIError{}
}

func (o RestAPIError) ToJson() string {
	js, serr := json.Marshal(o)
	if serr != nil {
		log.Println("Error while marshalling RestAPIError: ", serr)
	}
	return string(js)
}

func NewBadRequestError(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad request",
	}
}

func NewInternalServerError(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal server error",
	}
}
func NewNotFoundError(message string) RestAPIError {
	return RestAPIError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not found",
	}
}
