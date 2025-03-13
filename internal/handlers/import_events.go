package handlers

import (
	"calendar/internal/models"
	"calendar/internal/store"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"
)

// ImportEventsFromICS maneja la carga de un archivo .ics y lo guarda en la base de datos.
func ImportEventsFromICS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("📂 Recibiendo solicitud de importación...")

	if r.Method != http.MethodPost {
		fmt.Println("❌ Método no permitido")
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if store.DB == nil {
		fmt.Println("❌ Error: la base de datos no está inicializada")
		http.Error(w, "Error interno: la base de datos no está disponible", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("❌ Error al leer el archivo:", err)
		http.Error(w, "Error al leer el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Println("✅ Archivo recibido correctamente")

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("❌ Error al leer el contenido del archivo:", err)
		http.Error(w, "Error al leer el contenido del archivo", http.StatusInternalServerError)
		return
	}
	fmt.Println("✅ Contenido del archivo leído con éxito")

	// Parsear el contenido .ics usando un io.Reader.
	cal, err := ics.ParseCalendar(strings.NewReader(string(data)))
	if err != nil {
		fmt.Println("❌ Error al parsear el archivo .ics:", err)
		http.Error(w, "Error al parsear el archivo .ics", http.StatusInternalServerError)
		return
	}
	fmt.Println("✅ Archivo .ics parseado correctamente")

	var importedEvents []models.Event

	// Procesar cada evento del calendario.
	for _, event := range cal.Events() {
		fmt.Println("📅 Procesando evento:", event)

		// Verificar que existan las propiedades esenciales.
		summaryProp := event.GetProperty(ics.ComponentPropertySummary)
		if summaryProp == nil {
			fmt.Println("⚠️ Evento sin título, ignorado")
			continue
		}

		dtStartProp := event.GetProperty(ics.ComponentPropertyDtStart)
		dtEndProp := event.GetProperty(ics.ComponentPropertyDtEnd)
		if dtStartProp == nil || dtEndProp == nil {
			fmt.Println("⚠️ Evento sin fechas válidas, ignorado")
			continue
		}

		// Propiedades opcionales.
		var descriptionVal, locationVal string
		if descProp := event.GetProperty(ics.ComponentPropertyDescription); descProp != nil {
			descriptionVal = descProp.Value
		}
		if locProp := event.GetProperty(ics.ComponentPropertyLocation); locProp != nil {
			locationVal = locProp.Value
		}

		startTime, err := parseICalTime(dtStartProp.Value)
		if err != nil {
			fmt.Println("⚠️ Fecha de inicio inválida:", err)
			continue
		}

		endTime, err := parseICalTime(dtEndProp.Value)
		if err != nil {
			fmt.Println("⚠️ Fecha de fin inválida:", err)
			continue
		}

		newEvent := models.Event{
			Title:       summaryProp.Value,
			Description: descriptionVal,
			Location:    locationVal,
			StartTime:   startTime,
			EndTime:     endTime,
			IsAllDay:    isAllDayEvent(startTime, endTime),
		}

		if err := store.DB.Create(&newEvent).Error; err != nil {
			fmt.Println("❌ Error al guardar el evento:", err)
			continue
		}

		fmt.Println("✅ Evento guardado:", newEvent.Title)
		importedEvents = append(importedEvents, newEvent)
	}

	fmt.Println("✅ Importación finalizada con", len(importedEvents), "eventos")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(importedEvents)
}

// parseICalTime convierte una fecha de iCalendar a time.Time.
func parseICalTime(value string) (time.Time, error) {
	layouts := []string{"20060102T150405Z", "20060102T150405", "20060102"}
	for _, layout := range layouts {
		t, err := time.Parse(layout, value)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("error al parsear fecha: %s", value)
}

// isAllDayEvent verifica si un evento dura todo el día basándose en las horas de inicio y fin.
func isAllDayEvent(start, end time.Time) bool {
	return start.Hour() == 0 && start.Minute() == 0 && start.Second() == 0 &&
		end.Hour() == 0 && end.Minute() == 0 && end.Second() == 0
}
