package handlers

import (
	"calendar/internal/models"
	"encoding/json"
	"net/http"
	"strings"
)

// updateEvent maneja la actualización de un evento existente (PUT /events/{id})
func updateEvent(w http.ResponseWriter, r *http.Request) {
	// Extraer el ID del evento de la URL
	eventID := strings.TrimPrefix(r.URL.Path, "/events/")

	// Buscar el evento en la base de datos
	existingEvent, err := eventStore.GetEventByID(eventID)
	if err != nil {
		http.Error(w, "Evento no encontrado", http.StatusNotFound)
		return
	}

	// Decodificar los nuevos datos del evento
	var updatedEvent models.Event
	if err := json.NewDecoder(r.Body).Decode(&updatedEvent); err != nil {
		http.Error(w, "Error al decodificar el evento", http.StatusBadRequest)
		return
	}

	// Actualizar solo los campos que se proporcionan
	if updatedEvent.Title != "" {
		existingEvent.Title = updatedEvent.Title
	}
	if updatedEvent.Description != "" {
		existingEvent.Description = updatedEvent.Description
	}
	if !updatedEvent.StartTime.IsZero() {
		existingEvent.StartTime = updatedEvent.StartTime
	}
	if !updatedEvent.EndTime.IsZero() {
		existingEvent.EndTime = updatedEvent.EndTime
	}
	if updatedEvent.TimeZone != "" {
		existingEvent.TimeZone = updatedEvent.TimeZone
	}
	if updatedEvent.Location != "" {
		existingEvent.Location = updatedEvent.Location
	}
	existingEvent.IsAllDay = updatedEvent.IsAllDay
	if updatedEvent.Recurrence != "" {
		existingEvent.Recurrence = updatedEvent.Recurrence
	}

	// Validar el evento actualizado
	if err := existingEvent.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Guardar cambios en la base de datos
	if err := eventStore.UpdateEvent(existingEvent); err != nil {
		http.Error(w, "Error al actualizar el evento", http.StatusInternalServerError)
		return
	}

	// Responder con el evento actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingEvent)
}
