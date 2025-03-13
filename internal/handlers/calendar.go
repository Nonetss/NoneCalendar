package handlers

import (
	"calendar/internal/models"
	"calendar/internal/store"
	"calendar/pkg/icalendar"
	"encoding/json"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

var eventStore *store.EventStore

// Setup almacena la instancia de EventStore con la conexión DB
func SetupHandlers(db *gorm.DB) {
	eventStore = store.NewEventStore(db)
}

// HandleEvents maneja las solicitudes a /events
func HandleEvents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createEvent(w, r)
	case http.MethodGet:
		listEvents(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// HandleEventByID maneja las solicitudes a /events/{id}
func HandleEventByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getEventICS(w, r)
	case http.MethodDelete:
		deleteEvent(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// createEvent maneja la creación de un nuevo evento (POST /events)
func createEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Error al decodificar el evento", http.StatusBadRequest)
		return
	}

	// Validar el evento
	if err := event.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Almacenar el evento
	if err := eventStore.AddEvent(event); err != nil {
		http.Error(w, "Error al almacenar el evento", http.StatusInternalServerError)
		return
	}

	// Generar el archivo .ics
	icsContent := icalendar.GenerateICS(event)

	// Devolver el archivo .ics
	w.Header().Set("Content-Type", "text/calendar")
	w.Header().Set("Content-Disposition", "attachment; filename=event.ics")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(icsContent))
}

// listEvents maneja la obtención de todos los eventos (GET /events)
func listEvents(w http.ResponseWriter, r *http.Request) {
	events := eventStore.GetAllEvents()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

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

// deleteEvent maneja la eliminación de un evento (DELETE /events/{id})
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	// Extraer el ID del evento de la URL
	eventID := strings.TrimPrefix(r.URL.Path, "/events/")

	// Eliminar el evento
	if err := eventStore.DeleteEvent(eventID); err != nil {
		http.Error(w, "Evento no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Evento eliminado correctamente"))
}
