package icalendar

import (
	"calendar/internal/models"
	"fmt"
	"time"
)

// GenerateICS genera un archivo .ics para un solo evento
// (Se ignora la TZ y se usa UTC, o ajusta si deseas meter "TZID=Europe/Madrid").
func GenerateICS(event models.Event) string {
	dtStamp := time.Now().UTC().Format("20060102T150405Z")

	// Aquí, como tu código original, forzamos UTC. Si lo quieres con tz, ajusta:
	// dtStart := event.StartTime.UTC().Format("20060102T150405Z")
	// dtEnd   := event.EndTime.UTC().Format("20060102T150405Z")
	// O cambias a un TZ particular:

	dtStart := event.StartTime.Format("20060102T150405Z")
	dtEnd := event.EndTime.Format("20060102T150405Z")

	ics := "BEGIN:VCALENDAR\n"
	ics += "VERSION:2.0\n"
	ics += "PRODID:-//NoneLab//NoneCalendar//EN\n"
	ics += "BEGIN:VEVENT\n"
	ics += fmt.Sprintf("UID:%s@calendar.nonete.online\n", event.ID)
	ics += fmt.Sprintf("DTSTAMP:%s\n", dtStamp)
	ics += fmt.Sprintf("DTSTART:%s\n", dtStart)
	ics += fmt.Sprintf("DTEND:%s\n", dtEnd)
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
// todos en la zona horaria de Madrid.
func GenerateICSFeed(events []models.Event) string {
	loc, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		fmt.Println("No se pudo cargar la zona horaria Europe/Madrid. Se usará UTC.")
		loc = time.UTC
	}

	ics := "BEGIN:VCALENDAR\n"
	ics += "VERSION:2.0\n"
	ics += "PRODID:-//NoneLab//NoneCalendar//EN\n"

	// Definimos la VTIMEZONE una sola vez al inicio
	ics += "BEGIN:VTIMEZONE\n"
	ics += "TZID:Europe/Madrid\n"
	ics += "BEGIN:STANDARD\n"
	ics += "TZOFFSETFROM:+0200\n"
	ics += "TZOFFSETTO:+0100\n"
	ics += "TZNAME:CET\n"
	ics += "DTSTART:19701025T030000\n"
	ics += "END:STANDARD\n"
	ics += "BEGIN:DAYLIGHT\n"
	ics += "TZOFFSETFROM:+0100\n"
	ics += "TZOFFSETTO:+0200\n"
	ics += "TZNAME:CEST\n"
	ics += "DTSTART:19700329T020000\n"
	ics += "END:DAYLIGHT\n"
	ics += "END:VTIMEZONE\n"

	for _, event := range events {
		ics += "BEGIN:VEVENT\n"
		ics += fmt.Sprintf("UID:%s@calendar.nonete.online\n", event.ID)

		// El DTSTAMP es el momento de creación del evento (en UTC por convención).
		ics += fmt.Sprintf("DTSTAMP:%s\n", time.Now().UTC().Format("20060102T150405Z"))

		dtStartLocal := event.StartTime.In(loc).Format("20060102T150405")
		dtEndLocal := event.EndTime.In(loc).Format("20060102T150405")

		ics += fmt.Sprintf("DTSTART;TZID=Europe/Madrid:%s\n", dtStartLocal)
		ics += fmt.Sprintf("DTEND;TZID=Europe/Madrid:%s\n", dtEndLocal)

		ics += fmt.Sprintf("SUMMARY:%s\n", event.Title)

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
