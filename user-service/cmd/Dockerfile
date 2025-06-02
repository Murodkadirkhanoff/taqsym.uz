# 1. Этап сборки
FROM golang:1.24-alpine AS builder

# Установка зависимостей
RUN apk add --no-cache git

# Установка рабочей директории
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .
RUN go mod tidy
# Сборка бинарного файла
RUN go build -o bin/user-service ./user-service/cmd

# 2. Этап исполнения
FROM alpine:latest

# Установка сертификатов
RUN apk --no-cache add ca-certificates

# Создание папки для приложения
WORKDIR /root/

# Копируем бинарник из предыдущего этапа
COPY --from=builder /app/bin/user-service .

# Указываем команду запуска
CMD ["./user-service"]
