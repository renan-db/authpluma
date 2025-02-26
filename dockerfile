# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Instala dependências necessárias para compilação
RUN apk add --no-cache gcc musl-dev

# Instala o golang-migrate v4.18.2
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.2

# Copia os arquivos de dependência
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Compila a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

# Instala dependências necessárias
RUN apk add --no-cache ca-certificates tzdata

# Copia apenas o necessário
COPY --from=builder /app/main /app/main
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY ./internal/infrastructure/database/migration ./internal/infrastructure/database/migration

# Garante que o binário tenha permissão de execução
RUN chmod +x /app/main

# Usa a variável de ambiente para a porta
EXPOSE 8080

# Executa a aplicação
CMD ["/app/main"] 