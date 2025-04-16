FROM golang:1.19-alpine

# Установка таймзоны
RUN ln -snf /usr/share/zoneinfo/Asia/Almaty /etc/localtime && echo Asia/Almaty > /etc/timezone

# Установка зависимостей
RUN apk update && apk upgrade && apk add --no-cache git

# Рабочая директория
WORKDIR /app

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка бинарника
RUN CGO_ENABLED=0 GOOS=linux go build -o index

# Добавляем пользователя
RUN adduser -D -g 'app' app -u 1001 && chown -R app:app /app

# Меняем пользователя
USER app

# Запуск
ENTRYPOINT ["/app/index"]
