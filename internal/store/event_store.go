package store

import (
	"calendar/internal/models"
	"gorm.io/gorm"
)

// EventStore maneja los eventos con GORM
type EventStore struct {
	db *gorm.DB
}

// NewEventStore crea una instancia del almacenamiento de eventos usando GORM
func NewEventStore(db *gorm.DB) *EventStore {
	return &EventStore{db: db}
}

// AddEvent almacena un nuevo evento en la base de datos
func (es *EventStore) AddEvent(event models.Event) error {
	return es.db.Create(&event).Error
}

// GetEventByID busca un evento por su ID
func (es *EventStore) GetEventByID(id string) (*models.Event, error) {
	var event models.Event
	if err := es.db.First(&event, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

// GetAllEvents devuelve todos los eventos almacenados
func (es *EventStore) GetAllEvents() []models.Event {
	var events []models.Event
	es.db.Find(&events)
	return events
}

// DeleteEvent elimina un evento por su ID
func (es *EventStore) DeleteEvent(id string) error {
	return es.db.Delete(&models.Event{}, "id = ?", id).Error
}
