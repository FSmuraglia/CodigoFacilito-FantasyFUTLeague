# ---- Etapa 1: build ----
FROM golang:1.24-alpine AS builder

# Seteamos el directorio de trabajo
WORKDIR /app

# Copiamos los archivos de dependencias y descargamos módulos
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el resto del proyecto
COPY . .

# Compilamos el binario
RUN go build -o fantasyfutleague ./cmd/main.go

# ---- Etapa 2: runtime ----
FROM alpine:latest

# Directorio de trabajo en el contenedor final
WORKDIR /app

# Copiamos el binario desde la etapa anterior
COPY --from=builder /app/fantasyfutleague .

# Copiamos también los recursos estáticos y las plantillas
COPY static ./static
COPY templates ./templates

# Puerto en el que corre tu app (ajustalo si usás otro)
EXPOSE 8080

# Comando para ejecutar la app
CMD ["./fantasyfutleague"]
