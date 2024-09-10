#!/bin/bash

# Ожидание доступности базы данных
echo "Waiting for database to be ready..."
while ! pg_isready -h "postgresql" -p 5432 -U "postgres"; do
  sleep 2
done
echo "Database is ready!"

# Выполнение миграций
echo "Running database migrations..."
/app/migrate -path /app/migrations -database "postgresql://postgres:root@postgresql:5432/crypto-platform?sslmode=disable" -verbose up

if [ $? -ne 0 ]; then
  echo "Migrations failed!"
  exit 1
fi

# Запуск приложения
echo "Starting the application..."
exec ./app/app
