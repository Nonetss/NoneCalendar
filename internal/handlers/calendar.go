package handlers

import (
	"calendar/internal/store"
	"net/http"

	"gorm.io/gorm"
)

var eventStore *store.EventStore

// SetupHandlers inicializa los handlers con la conexión a la base de datos
func SetupHandlers(db *gorm.DB) {
	eventStore = store.NewEventStore(db)

	// Rutas del API
	http.HandleFunc("/events", HandleEvents)
	http.HandleFunc("/events/", HandleEventByID)
	http.HandleFunc("/calendar.ics", GetAllEventsICS)
}

// HandleEvents maneja las solicitudes a /events
func HandleEvents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createEvent(w, r)
	case http.MethodGet:
		listEvents(w, r)
	case http.MethodPut:
		updateEvent(w, r)
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
	case http.MethodPut:
		updateEvent(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
