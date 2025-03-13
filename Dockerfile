# Etapa 1: Compilar la aplicación
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Copiar los archivos de módulos y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo el código fuente
COPY . .

# Compilar el binario. Se usa CGO deshabilitado para mayor portabilidad.
RUN CGO_ENABLED=0 go build -o calendar ./cmd/api

# Etapa 2: Imagen final minimalista
FROM alpine:latest
WORKDIR /app

# Copiar el binario compilado
COPY --from=builder /app/calendar .

# Opcional: copiar el archivo .env si lo necesitas en la imagen (o usa variables de entorno en docker-compose)
COPY .env .

# Exponer el puerto en el que corre la API
EXPOSE 8080

# Ejecutar la aplicación
ENTRYPOINT ["./calendar"]
