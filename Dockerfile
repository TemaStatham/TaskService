# Stage 1: Build
FROM golang:latest AS build

WORKDIR /app

# Копируем зависимости и устанавливаем их
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp ./cmd

# Stage 2: Run
FROM alpine:latest AS run

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарник
COPY --from=build /myapp /myapp

COPY config.yaml ./config.yaml

# Разрешаем выполнение файла (на всякий случай)
RUN chmod +x /myapp

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["/myapp", "--config=./app/config.yaml"]
