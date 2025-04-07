# Используем официальный образ Golang
FROM golang:latest AS builder

WORKDIR /app

# Устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Компилируем приложение
RUN go build -o films_service ./cmd/main.go


# Финальный образ
FROM alpine:latest

WORKDIR /app

# Устанавливаем необходимые пакеты (если используем Alpine)
RUN apk add --no-cache bash ca-certificates

# Копируем бинарник приложения
COPY --from=builder /app/movie_service .

FROM gomicro/goose AS final

WORKDIR /migrations

# Копируем миграции и entrypoint.sh
COPY ./migrations /migrations
COPY entrypoint.sh /migrations/entrypoint.sh

RUN chmod +x /migrations/entrypoint.sh

ENTRYPOINT ["/migrations/entrypoint.sh"]