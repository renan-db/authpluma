# Etapa de construção
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /authpluma ./cmd/api

# Etapa final
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /authpluma .

EXPOSE 8080

CMD ["./authpluma"] 