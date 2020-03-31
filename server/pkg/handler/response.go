package handler

import (
	"encoding/json"
	"net/http"
)

// ErrorPayload is the error object that gets returned in the event of an error.
type ErrorPayload struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ResponsePayload is the payload that gets returned, which includes data or error
type ResponsePayload struct {
	Data  interface{}   `json:"data"`
	Error *ErrorPayload `json:"error"`
}

// RespondWithSuccess is a helper function to respond in a HTTP-friendly way to be used
// throughout the app following RESTful standards
func RespondWithSuccess(w http.ResponseWriter, statusCode int, payload interface{}) {
	responsePayload := ResponsePayload{
		Data:  payload,
		Error: nil,
	}

	response, err := json.Marshal(responsePayload)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error marshaling the response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	_, err = w.Write(response)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error writing the response")
		return
	}
}

// RespondWithError is a helper function to respond in a HTTP-friendly way to be used
// throughout the app following RESTful standards
func RespondWithError(w http.ResponseWriter, statusCode int, errMsg string) {
	responsePayload := ResponsePayload{
		Data: nil,
		Error: &ErrorPayload{
			Code:    statusCode,
			Message: errMsg,
		},
	}

	response, err := json.Marshal(responsePayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(statusCode)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
