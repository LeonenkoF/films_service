#!/bin/bash
#драйвер библиотеки миграций
GOOSE_DRIVER=postgres
#режим подключения к БД
SSL_MODE=disable
#путь к миграциям
GOOSE_MIGRATION_DIR=./app/migrations 
#тип СУБД
DB_TYPE=postgres
DB_HOST=postgres_movie_service

DBSTRING="host=$DB_HOST user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=$SSL_MODE"

goose postgres "$DBSTRING" up