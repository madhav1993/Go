package handlers

import (
	"encoding/json"
	"net/http"
)

func readiness(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "Ok",
	}

	json.NewEncoder(w).Encode(status)
}
