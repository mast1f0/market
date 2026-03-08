include configs/.env

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATIONS_PATH=./migrations

.PHONY: run migrate-up migrate-down migrate-reset

run:
	cd cmd/srv && air
migrate-up:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) up

migrate-down:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) down

migrate-reset:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) drop -f
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) up