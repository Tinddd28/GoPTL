# Используем официальный образ Golang для сборки
FROM golang:1.23

# Устанавливаем переменную GOPATH
ENV GOPATH=/

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь остальной код
COPY . .

# Делаем скрипт wait-for-postgres.sh исполняемым (если он нужен)
RUN chmod +x wait-for-postgres.sh

# Сборка приложения
RUN go build -o goptl ./cmd/main.go

# Указываем команду по умолчанию для запуска приложения
CMD ["./goptl"]
