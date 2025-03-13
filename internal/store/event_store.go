package store

import (
	"calendar/internal/models"
	"errors"
	"sync"
)

// EventStore almacena los eventos en memoria.
type EventStore struct {
	events map[string]models.Event // Mapa para acceder rápidamente a los eventos por ID
	mu     sync.RWMutex            // Mutex para garantizar la seguridad en concurrencia
}

// NewEventStore crea una nueva instancia de EventStore.
func NewEventStore() *EventStore {
	return &EventStore{
		events: make(map[string]models.Event),
	}
}

// AddEvent agrega un nuevo evento al almacenamiento.
func (es *EventStore) AddEvent(event models.Event) error {
	es.mu.Lock()
	defer es.mu.Unlock()

	// Verificar si el evento ya existe
	if _, exists := es.events[event.ID]; exists {
		return errors.New("el evento ya existe")
	}

	// Agregar el evento al mapa
	es.events[event.ID] = event
	return nil
}

// GetEventByID busca un evento por su ID.
func (es *EventStore) GetEventByID(id string) (*models.Event, error) {
	es.mu.RLock()
	defer es.mu.RUnlock()

	// Buscar el evento en el mapa
	event, exists := es.events[id]
	if !exists {
		return nil, errors.New("evento no encontrado")
	}

	return &event, nil
}

// GetAllEvents devuelve todos los eventos almacenados.
func (es *EventStore) GetAllEvents() []models.Event {
	es.mu.RLock()
	defer es.mu.RUnlock()

	// Crear una lista de eventos
	events := make([]models.Event, 0, len(es.events))
	for _, event := range es.events {
		events = append(events, event)
	}

	return events
}

// DeleteEvent elimina un evento por su ID.
func (es *EventStore) DeleteEvent(id string) error {
	es.mu.Lock()
	defer es.mu.Unlock()

	// Verificar si el evento existe
	if _, exists := es.events[id]; !exists {
		return errors.New("evento no encontrado")
	}

	// Eliminar el evento del mapa
	delete(es.events, id)
	return nil
}

