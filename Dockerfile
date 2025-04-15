# Используем официальный образ Golang
FROM golang:latest AS builder

WORKDIR /app

# Устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Компилируем приложение
RUN go build -o main ./main.go


FROM gomicro/goose AS final

COPY --from=builder /app/main .

# Копируем миграции и entrypoint.sh
COPY ./migrations /migrations
COPY entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/entrypoint.sh

WORKDIR /app

ENTRYPOINT ["/app/entrypoint.sh"]

CMD ["./main"]