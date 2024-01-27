# Dockerfile
FROM golang:1.21

WORKDIR /app

# Copiar el go.mod para inicializar el módulo
COPY go.mod .

# Instalar dependencias y herramientas
RUN go mod tidy
RUN go mod download

# Copiar el resto del código
COPY . .

# Build de la aplicación
RUN go build -o command-app cmd/command/main.go

# Comando para ejecutar las migraciones
CMD ["./command-app"]
