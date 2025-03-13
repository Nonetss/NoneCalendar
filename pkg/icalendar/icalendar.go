package icalendar

import (
	"calendar/internal/models"
	"fmt"
	"time"
)

// GenerateICS genera un archivo .ics para un solo evento
func GenerateICS(event models.Event) string {
	// Formatear la fecha de sello de tiempo
	dtStamp := time.Now().UTC().Format("20060102T150405Z")

	// Formatear las fechas del evento
	var dtStart, dtEnd string
	if event.IsAllDay {
		// Para eventos de todo el día, DTEND debe ser el día siguiente
		dtStart = event.StartTime.UTC().Format("20060102")
		dtEnd = event.EndTime.UTC().Add(24 * time.Hour).Format("20060102")
	} else {
		if event.TimeZone != "" {
			// Si hay zona horaria, usar TZID
			dtStart = fmt.Sprintf("TZID=%s:%s", event.TimeZone, event.StartTime.Format("20060102T150405"))
			dtEnd = fmt.Sprintf("TZID=%s:%s", event.TimeZone, event.EndTime.Format("20060102T150405"))
		} else {
			// Si no hay zona horaria, usar UTC
			dtStart = event.StartTime.UTC().Format("20060102T150405Z")
			dtEnd = event.EndTime.UTC().Format("20060102T150405Z")
		}
	}

	// Generar contenido .ics
	ics := "BEGIN:VCALENDAR\n"
	ics += "VERSION:2.0\n"
	ics += "PRODID:-//NoneLab//NoneCalendar//EN\n"
	ics += "BEGIN:VEVENT\n"
	ics += fmt.Sprintf("UID:%s@calendar.nonete.online\n", event.ID)
	ics += fmt.Sprintf("DTSTAMP:%s\n", dtStamp)

	if event.IsAllDay {
		ics += fmt.Sprintf("DTSTART;VALUE=DATE:%s\n", dtStart)
		ics += fmt.Sprintf("DTEND;VALUE=DATE:%s\n", dtEnd)
	} else {
		ics += fmt.Sprintf("DTSTART:%s\n", dtStart)
		ics += fmt.Sprintf("DTEND:%s\n", dtEnd)
	}

	ics += fmt.Sprintf("SUMMARY:%s\n", event.Title)

	if event.Description != "" {
		ics += fmt.Sprintf("DESCRIPTION:%s\n", event.Description)
	}
	if event.Location != "" {
		ics += fmt.Sprintf("LOCATION:%s\n", event.Location)
	}

	ics += "END:VEVENT\n"
	ics += "END:VCALENDAR\n"

	return ics
}

// GenerateICSFeed genera un archivo .ics con múltiples eventos
func GenerateICSFeed(events []models.Event) string {
	ics := "BEGIN:VCALENDAR\n"
	ics += "VERSION:2.0\n"
	ics += "PRODID:-//NoneLab//NoneCalendar//EN\n"

	for _, event := range events {
		ics += "BEGIN:VEVENT\n"
		ics += fmt.Sprintf("UID:%s@calendar.nonete.online\n", event.ID)
		ics += fmt.Sprintf("SUMMARY:%s\n", event.Title)
		ics += fmt.Sprintf("DTSTAMP:%s\n", time.Now().UTC().Format("20060102T150405Z"))

		var dtStart, dtEnd string
		if event.IsAllDay {
			dtStart = event.StartTime.UTC().Format("20060102")
			dtEnd = event.EndTime.UTC().Add(24 * time.Hour).Format("20060102")
			ics += fmt.Sprintf("DTSTART;VALUE=DATE:%s\n", dtStart)
			ics += fmt.Sprintf("DTEND;VALUE=DATE:%s\n", dtEnd)
		} else {
			if event.TimeZone != "" {
				dtStart = fmt.Sprintf("TZID=%s:%s", event.TimeZone, event.StartTime.Format("20060102T150405"))
				dtEnd = fmt.Sprintf("TZID=%s:%s", event.TimeZone, event.EndTime.Format("20060102T150405"))
			} else {
				dtStart = event.StartTime.UTC().Format("20060102T150405Z")
				dtEnd = event.EndTime.UTC().Format("20060102T150405Z")
			}
			ics += fmt.Sprintf("DTSTART:%s\n", dtStart)
			ics += fmt.Sprintf("DTEND:%s\n", dtEnd)
		}

		if event.Description != "" {
			ics += fmt.Sprintf("DESCRIPTION:%s\n", event.Description)
		}
		if event.Location != "" {
			ics += fmt.Sprintf("LOCATION:%s\n", event.Location)
		}

		ics += "END:VEVENT\n"
	}

	ics += "END:VCALENDAR\n"
	return ics
}
