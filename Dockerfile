# Сборочный этап
FROM golang:1.21-alpine AS build

WORKDIR /app
COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o wallet-service ./cmd/main.go

# Финальный контейнер
FROM alpine:latest

WORKDIR /root/
COPY --from=build /app/wallet-service .

EXPOSE 8080

CMD ["./wallet-service"]