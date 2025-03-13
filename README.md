# Estructura de carpetas dentro del proyecto:

```
/calendar
│
├── /cmd
│   └── /api
│       └── main.go          # Punto de entrada de la aplicación
│
├── /internal
│   ├── /handlers            # Manejadores de las rutas HTTP
│   │   └── calendar.go
│   ├── /models              # Estructuras de datos (como tu struct Event)
│   │   └── event.go
│   ├── /services            # Lógica de negocio (generación de .ics, etc.)
│   │   └── ics_service.go
│   └── /utils               # Utilidades comunes (manejo de fechas, validaciones, etc.)
│       └── date_utils.go
│
├── /pkg
│   └── /icalendar            # Paquete para generar y manipular archivos .ics
│       └── icalendar.go
│
├── /api
│   └── /docs                 # Documentación de la API (Swagger, OpenAPI, etc.)
│       └── swagger.yaml
│
├── /configs                  # Archivos de configuración (JSON, YAML, etc.)
│   └── config.yaml
│
├── /scripts                  # Scripts útiles (migraciones, despliegues, etc.)
│   └── deploy.sh
│
├── /web                      # Archivos estáticos (si es necesario)
│   └── /static
│       └── index.html
│
├── go.mod                    # Archivo de módulos de Go
├── go.sum                    # Archivo de sumas de verificación de Go
└── README.md                 # Documentación del proyecto
```

---

# Endpoint API:

1. **`POST /events`**:

   - **Descripción**: Crea un nuevo evento.
   - **Entrada**: Un JSON con la estructura de `Event`.
   - **Salida**: Un archivo `.ics` que representa el evento creado.
   - **Uso**: El cliente envía los datos del evento y recibe el archivo `.ics` para descargar.

2. **`GET /events/{id}.ics`**:

   - **Descripción**: Obtiene un archivo `.ics` para un evento específico.
   - **Entrada**: El `id` del evento.
   - **Salida**: Un archivo `.ics` que representa el evento.
   - **Uso**: El cliente puede descargar el archivo `.ics` para importarlo en Google Calendar.

3. **`GET /events`**:

   - **Descripción**: Obtiene una lista de todos los eventos.
   - **Entrada**: Ninguna.
   - **Salida**: Un JSON con la lista de eventos.
   - **Uso**: Para ver todos los eventos creados.

4. **`DELETE /events/{id}`**:
   - **Descripción**: Elimina un evento específico.
   - **Entrada**: El `id` del evento.
   - **Salida**: Un mensaje de confirmación.
   - **Uso**: Para eliminar eventos que ya no son necesarios.
