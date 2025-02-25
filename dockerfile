# Build stage
FROM golang:1.22 AS builder

WORKDIR /app

# Copia os arquivos de dependência
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Compila a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

# Adiciona certificados CA e timezone
RUN apk --no-cache add ca-certificates tzdata

# Copia o binário compilado
COPY --from=builder /app/main .

# Usa a variável de ambiente para a porta
EXPOSE ${APP_PORT}

# Executa a aplicação
CMD ["./main"] 