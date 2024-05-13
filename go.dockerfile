# # Этап 1: сборка приложения и миграции
# FROM golang:latest AS builder

# WORKDIR /app

# # Копируем go.mod и go.sum и загружаем зависимости
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# # Копируем все файлы проекта в текущую директорию
# COPY . .

# # Собираем бинарный файл приложения
# RUN go build -o main ./cmd/.

# # Этап 2: запуск миграций и копирование бинарного файла
# FROM builder AS migrator

# # Копируем миграции внутрь образа
# COPY ./db/migrations /migrations

# # Выполняем миграции
# RUN migrate -database postgres://postgres:postgres@db:5432/friendlorant?sslmode=disable -path /migrations up

# # Этап 3: создание окончательного образа
# FROM postgres:16

# # Копируем бинарный файл приложения из этапа 1
# COPY --from=builder /app/main /app/main

# # Устанавливаем переменную окружения для указания рабочей директории
# WORKDIR /app

# # Устанавливаем переменные окружения для базы данных
# ENV POSTGRES_USER=postgres
# ENV POSTGRES_PASSWORD=postgres
# ENV POSTGRES_DB=friendlorant
# ENV DATABASE_URL=postgres://postgres:postgres@db:5432/friendlorant?sslmode=disable

# # Устанавливаем порт, который будет использоваться приложением
# EXPOSE 8000

# # Запускаем приложение
# CMD ["./main"]
