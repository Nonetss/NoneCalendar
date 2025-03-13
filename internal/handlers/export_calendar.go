package handlers

import (
	"calendar/pkg/icalendar"
	"net/http"
)

// GetAllEventsICS genera un archivo .ics con todos los eventos (GET /calendar.ics)
func GetAllEventsICS(w http.ResponseWriter, r *http.Request) {
	// Obtener todos los eventos desde la base de datos
	events := eventStore.GetAllEvents()

	// Generar el contenido .ics para todos los eventos
	icsContent := icalendar.GenerateICSFeed(events)

	// Configurar cabeceras HTTP para indicar que es un archivo .ics
	w.Header().Set("Content-Type", "text/calendar")
	w.Header().Set("Content-Disposition", "attachment; filename=calendar.ics")
	w.Write([]byte(icsContent))
}
