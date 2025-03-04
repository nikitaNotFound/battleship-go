-include .env

MIGRATIONS_DIR := internal/storage/postgres/migrations
DB_URL := "postgres://$(DB_POSTGRES_USER):$(DB_POSTGRES_PASSWORD)@$(DB_POSTGRES_HOST):$(DB_POSTGRES_PORT)/$(DB_POSTGRES_NAME)"

.PHONY: migrate-up
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

.PHONY: migrate-status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" status

.PHONY: migrate-create
migrate-create:
	@read -p "Enter migration name: " name; \
	goose -dir $(MIGRATIONS_DIR) create $${name} sql

.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: docker-build
docker-build:
	docker build -t battleship:latest .

.PHONY: docker-compose-up
docker-compose-up:
	docker-compose -f deployments/docker-compose.yml up -d

.PHONY: docker-compose-down
docker-compose-down:
	docker-compose -f deployments/docker-compose.yml down

.PHONY: docker-run
docker-run: docker-build docker-compose-up