package handlers

import (
	"encoding/json"
	"net/http"
)

// listEvents maneja la obtención de todos los eventos (GET /events)
func listEvents(w http.ResponseWriter, r *http.Request) {
	events := eventStore.GetAllEvents()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
