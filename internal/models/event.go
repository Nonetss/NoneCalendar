package main

import (
	"errors"
	"time"
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
