package server

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"health": "ok"}`))
	if err != nil {
		return
	}
}
