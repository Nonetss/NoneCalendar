package models

import (
	"testing"
	"time"
)

func TestEventValidation(t *testing.T) {
	event := Event{
		ID:        "1",
		Title:     "Evento de prueba",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
	}

	if err := event.Validate(); err != nil {
		t.Errorf("Error validando evento: %v", err)
	}
}

func TestEventStore(t *testing.T) {
	store := NewEventStore()

	event := Event{
		ID:        "1",
		Title:     "Evento de prueba",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
	}

	if err := store.AddEvent(event); err != nil {
		t.Errorf("Error agregando evento: %v", err)
	}

	if len(store.GetAllEvents()) != 1 {
		t.Errorf("Se esperaba 1 evento, se encontraron %d", len(store.GetAllEvents()))
	}

	foundEvent, err := store.GetEventByID("1")
	if err != nil {
		t.Errorf("Error buscando evento: %v", err)
	}
	if foundEvent.Title != "Evento de prueba" {
		t.Errorf("El evento encontrado no coincide")
	}
}
