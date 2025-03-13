package icalendar

import (
	"calendar/internal/models"
	"fmt"
	"time"
)

// GenerateICS genera un archivo .ics a partir de un evento.
func GenerateICS(event models.Event) string {
	// Formatear las fechas en el formato requerido por iCalendar
	dtStamp := time.Now().UTC().Format("20060102T150405Z")

	var dtStart, dtEnd string
	if event.IsAllDay {
		// Formato para eventos de todo el día
		dtStart = event.StartTime.UTC().Format("20060102")
		dtEnd = event.EndTime.UTC().Format("20060102")
	} else {
		// Formato para eventos con hora específica
		if event.TimeZone != "" {
			// Usar TZID para indicar la zona horaria
			dtStart = fmt.Sprintf("TZID=%s:%s", event.TimeZone, event.StartTime.Format("20060102T150405"))
			dtEnd = fmt.Sprintf("TZID=%s:%s", event.TimeZone, event.EndTime.Format("20060102T150405"))
		} else {
			// Si no hay zona horaria, usar UTC
			dtStart = event.StartTime.UTC().Format("20060102T150405Z")
			dtEnd = event.EndTime.UTC().Format("20060102T150405Z")
		}
	}

	// Crear el contenido del archivo .ics
	ics := "BEGIN:VCALENDAR\n"
	ics += "VERSION:2.0\n"
	ics += "PRODID:-//Nonelab//NoneCalendar//EN\n"
	ics += "BEGIN:VEVENT\n"
	ics += fmt.Sprintf("UID:%s@nonecalendar\n", event.ID)
	ics += fmt.Sprintf("DTSTAMP:%s\n", dtStamp)

	if event.IsAllDay {
		ics += fmt.Sprintf("DTSTART;VALUE=DATE:%s\n", dtStart)
		ics += fmt.Sprintf("DTEND;VALUE=DATE:%s\n", dtEnd)
	} else {
		ics += fmt.Sprintf("DTSTART;%s\n", dtStart)
		ics += fmt.Sprintf("DTEND;%s\n", dtEnd)
	}

	ics += fmt.Sprintf("SUMMARY:%s\n", event.Title)
	ics += fmt.Sprintf("DESCRIPTION:%s\n", event.Description)
	ics += fmt.Sprintf("LOCATION:%s\n", event.Location)
	ics += "END:VEVENT\n"
	ics += "END:VCALENDAR\n"

	return ics
}
