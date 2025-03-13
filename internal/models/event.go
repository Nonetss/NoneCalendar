package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Event representa un evento en el calendario.
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	TimeZone    string    `json:"time_zone"`
	Location    string    `json:"location"`
	IsAllDay    bool      `json:"is_all_day"`
	Recurrence  string    `json:"recurrence"`
}

func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == "" {
		e.ID = uuid.New().String()[:8] // Generar un ID más corto
	}
	return
}

// Validate valida que los datos del evento sean correctos.
func (e *Event) Validate() error {
	if e.Title == "" {
		return errors.New("el título del evento es obligatorio")
	}
	if e.StartTime.IsZero() {
		return errors.New("la fecha de inicio es obligatoria")
	}
	if e.EndTime.IsZero() {
		return errors.New("la fecha de fin es obligatoria")
	}
	if e.StartTime.After(e.EndTime) {
		return errors.New("la fecha de inicio debe ser anterior a la fecha de fin")
	}
	return nil
}

// EventStore almacena los eventos en memoria.
type EventStore struct {
	events []Event
}

// NewEventStore crea una nueva instancia de EventStore.
func NewEventStore() *EventStore {
	return &EventStore{
		events: []Event{},
	}
}

// AddEvent agrega un nuevo evento al almacenamiento.
func (es *EventStore) AddEvent(event Event) error {
	if err := event.Validate(); err != nil {
		return err
	}
	es.events = append(es.events, event)
	return nil
}

// GetEventByID busca un evento por su ID.
func (es *EventStore) GetEventByID(id string) (*Event, error) {
	for _, event := range es.events {
		if event.ID == id {
			return &event, nil
		}
	}
	return nil, errors.New("evento no encontrado")
}

// GetAllEvents devuelve todos los eventos almacenados.
func (es *EventStore) GetAllEvents() []Event {
	return es.events
}
