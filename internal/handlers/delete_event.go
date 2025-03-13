package handlers

import (
	"net/http"
	"strings"
)

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
