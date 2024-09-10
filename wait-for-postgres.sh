#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

# Переменные окружения для подключения к базе данных
PGHOST="$host"
PGPORT="${DB_PORT:-5432}"
PGUSER="${DB_USER:-postgres}"
PGPASSWORD="${DB_PASSWORD:-root}"
DBNAME="${DB_NAME:-crypto-platform}"

# Ожидание, пока база данных станет доступной
until psql -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - creating database if not exists"

# Создание базы данных, если она не существует
psql -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" <<-EOSQL
  CREATE DATABASE IF NOT EXISTS $DBNAME;
EOSQL

>&2 echo "Postgres is up - executing command"
exec $cmd
