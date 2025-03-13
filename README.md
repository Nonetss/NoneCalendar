# Estructura del proyecto:

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
