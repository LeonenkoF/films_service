#драйвер библиотеки миграций
GOOSE_DRIVER=postgres
#режим подключения к БД
SSL_MODE=disable
#путь к миграциям
GOOSE_MIGRATION_DIR=./migrations 
#тип СУБД
DB_TYPE=postgres
DB_HOST=localhost
#строка подключения к БД
GOOSE_DBSTRING=$(DB_TYPE)://$(DB_USER):"qwerty"@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)

# migrate-up применяет все схемы миграций
migrate-up:
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	goose -dir $(GOOSE_MIGRATION_DIR) \
	up

# migrate-down откатывает последнюю схему миграции
migrate-down:
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	goose -dir $(GOOSE_MIGRATION_DIR) \
	down

# migrate-status проверяет статус схемы миграций
migrate-status:
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	goose -dir $(GOOSE_MIGRATION_DIR) \
	status

# migrate-create создает новый файл миграции с заданным именем
migrate-create:
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	goose -dir $(GOOSE_MIGRATION_DIR) \
	create $(name) sql

# build создаёт образ по docker-compose-файлу
build:
	docker-compose build 

# run-local запускает контейнеры по docker-compose-файлу
run-local:
	docker-compose up