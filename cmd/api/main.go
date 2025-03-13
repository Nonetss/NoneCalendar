package main

import (
	"calendar/internal/handlers"
	"calendar/internal/models"
	"calendar/internal/store"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Iniciar la conexión con la base de datos
	store.InitDB()

	fmt.Println("🟡 Ejecutando migración de la base de datos...")

	// 🔹 FORZAR MIGRACIÓN 🔹
	if err := store.DB.AutoMigrate(&models.Event{}); err != nil {
		log.Fatalf("❌ Error en la migración de la base de datos: %v", err)
	}

	fmt.Println("✅ Migración completada. La tabla 'events' debería existir ahora.")

	// Configurar handlers con la conexión a la base de datos
	handlers.SetupHandlers(store.DB)

	// Iniciar el servidor
	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
