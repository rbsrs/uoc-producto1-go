# --- Etapa 1: build (compilación) ---
FROM golang:1.22 AS builder
WORKDIR /app

# Copiamos metadatos del módulo primero (mejor caché)
COPY go.mod ./
RUN go mod download

# Copiamos el resto del proyecto
COPY . .

# Compilamos binario Linux. CGO_DISABLED evita dependencias nativas
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# --- Etapa 2: runtime (ejecución) ---
# Opción A (recomendada): distroless (mínima, sin shell)
FROM gcr.io/distroless/base-debian12
WORKDIR /app

COPY --from=builder /app/server /app/server
COPY --from=builder /app/static /app/static

EXPOSE 8080
USER nonroot:nonroot
ENV PORT=8080
ENTRYPOINT ["/app/server"]