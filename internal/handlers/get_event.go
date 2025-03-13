package handlers

import (
	"calendar/pkg/icalendar"
	"net/http"
	"strings"
)

// getEventICS maneja la obtención de un archivo .ics para un evento específico (GET /events/{id}.ics)
func getEventICS(w http.ResponseWriter, r *http.Request) {
	// Extraer el ID del evento de la URL
	eventID := strings.TrimPrefix(r.URL.Path, "/events/")
	eventID = strings.TrimSuffix(eventID, ".ics")

	// Buscar el evento
	event, err := eventStore.GetEventByID(eventID)
	if err != nil {
		http.Error(w, "Evento no encontrado", http.StatusNotFound)
		return
	}

	// Generar el archivo .ics
	icsContent := icalendar.GenerateICS(*event)

	// Devolver el archivo .ics
	w.Header().Set("Content-Type", "text/calendar")
	w.Header().Set("Content-Disposition", "attachment; filename=event.ics")
	w.Write([]byte(icsContent))
}
