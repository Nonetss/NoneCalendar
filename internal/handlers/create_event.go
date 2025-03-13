package handlers

import (
	"calendar/internal/models"
	"calendar/pkg/icalendar"
	"encoding/json"
	"net/http"
)

// createEvent maneja la creación de un nuevo evento (POST /events)
func createEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Error al decodificar el evento", http.StatusBadRequest)
		return
	}

	// Generar ID automáticamente si no viene en la petición
	event.ID = ""

	// Validar el evento
	if err := event.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Almacenar el evento en la base de datos
	if err := eventStore.AddEvent(event); err != nil {
		http.Error(w, "Error al almacenar el evento", http.StatusInternalServerError)
		return
	}

	// Generar el archivo .ics (sin el segundo parámetro)
	icsContent := icalendar.GenerateICS(event)

	// Devolver el archivo .ics
	w.Header().Set("Content-Type", "text/calendar")
	w.Header().Set("Content-Disposition", "attachment; filename=event.ics")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(icsContent))
}
