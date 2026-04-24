-include configs/.env

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATIONS_PATH=./migrations

.PHONY: run migrate-up migrate-down migrate-reset env-generate create-db

run-dev:
	go run cmd/srv/main.go
migrate-up:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) up

migrate-down:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) down

migrate-reset:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) drop -f
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) up

env-generate:
	echo "Generating default config !!!CHANGE PARAMS"
	if [ ! -d "configs" ]; then mkdir configs; fi
	if [ ! -f "configs/.env" ]; then cd configs && touch .env; fi
	echo -e "DB_USER=store_usr\n\
DB_PASSWORD=password132\n\
DB_HOST=localhost\n\
DB_PORT=5432\n\
DB_NAME=store_db\n\
DB_SSLMODE=disable\n" > configs/.env

# проверка юзера не работает
create-db:
	@psql -h $(DB_HOST) -p $(DB_PORT) -U postgres -tc "SELECT 1 FROM pg_roles WHERE rolname='$(DB_USER)'" | grep -q 1 || \
	psql -h $(DB_HOST) -p $(DB_PORT) -U postgres -c "CREATE USER $(DB_USER) WITH PASSWORD '$(DB_PASSWORD)';"

	@psql -h $(DB_HOST) -p $(DB_PORT) -U postgres -tc "SELECT 1 FROM pg_database WHERE datname='$(DB_NAME)'" | grep -q 1 || \
	psql -h $(DB_HOST) -p $(DB_PORT) -U postgres -c "CREATE DATABASE $(DB_NAME) OWNER $(DB_USER);"

	@psql -h $(DB_HOST) -p $(DB_PORT) -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE $(DB_NAME) TO $(DB_USER);"