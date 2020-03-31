package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ResponseWithSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	payload := "test-payload"
	statusCode := http.StatusOK
	RespondWithSuccess(w, statusCode, payload)

	desiredPayload := ResponsePayload{
		Data:  payload,
		Error: nil,
	}
	bytes, err := json.Marshal(desiredPayload)
	if err != nil {
		t.Errorf("Error marshaling desired payload: %w", err)
	}

	if w.Body.String() != string(bytes) {
		t.Errorf("actual payload doesn't match expected. \nGot: %s, \nWant: %s", w.Body.String(), string(bytes))
	}
}

func Test_ResponseWithError(t *testing.T) {
	w := httptest.NewRecorder()

	errMsg := "error-msg"
	statusCode := http.StatusInternalServerError
	RespondWithError(w, statusCode, errMsg)

	desiredPayload := ResponsePayload{
		Data: nil,
		Error: &ErrorPayload{
			Code:    statusCode,
			Message: errMsg,
		},
	}
	bytes, err := json.Marshal(desiredPayload)
	if err != nil {
		t.Errorf("Error marshaling desired payload: %w", err)
	}

	if w.Body.String() != string(bytes) {
		t.Errorf("actual payload doesn't match expected. \nGot: %s, \nWant: %s", w.Body.String(), string(bytes))
	}
}
