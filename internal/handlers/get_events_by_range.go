package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

// getEventsByDateRange maneja la obtención de eventos dentro de un rango de fechas (GET /events/range?start=YYYY-MM-DD&end=YYYY-MM-DD)
func getEventsByDateRange(w http.ResponseWriter, r *http.Request) {
	// Obtener los parámetros `start` y `end` de la URL
	startDateStr := r.URL.Query().Get("start")
	endDateStr := r.URL.Query().Get("end")

	// Validar que ambos parámetros fueron proporcionados
	if startDateStr == "" || endDateStr == "" {
		http.Error(w, "Debe proporcionar los parámetros 'start' y 'end'", http.StatusBadRequest)
		return
	}

	// Convertir las fechas a formato `time.Time`
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		http.Error(w, "Formato de fecha 'start' inválido. Use YYYY-MM-DD.", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		http.Error(w, "Formato de fecha 'end' inválido. Use YYYY-MM-DD.", http.StatusBadRequest)
		return
	}

	// Obtener los eventos dentro del rango
	events := eventStore.GetEventsByDateRange(startDate, endDate)

	// Devolver la lista de eventos en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
