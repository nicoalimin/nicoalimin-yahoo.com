package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HealthHandler_Success(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handlerHealth := Health{}
	handlerHealth.HealthCheck(w, r)

	if w.Code != http.StatusOK {
		t.Errorf(
			"expected status ok when calling health handler endpoint. \nGot: %d \nWant: %d",
			w.Code,
			http.StatusOK,
		)
	}
}
