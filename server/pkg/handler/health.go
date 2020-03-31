package handler

import (
	"net/http"
)

// Health contains methods that is related to the healthiness of the server
type Health struct{}

// HealthCheck responds with a string containing a friendly message indicating
// that the server is up and running
func (h *Health) HealthCheck(w http.ResponseWriter, r *http.Request) {
	RespondWithSuccess(w, http.StatusOK, "I am ok, everything is fine, so fine.")
}
