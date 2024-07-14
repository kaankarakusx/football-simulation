.PHONY: migrate-up migrate-down  build run test

MIGRATE_CMD=go run ./cmd/migrate/main.go
MIGRATE_DIR=./cmd/migrate/migrations
MAIN_PACKAGE=./cmd/main.go

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATE_DIR) $$name

migrate-up:
	$(MIGRATE_CMD) up

migrate-down:
	$(MIGRATE_CMD) down

build:
	go build -o bin/football-simulation $(MAIN_PACKAGE)

run:
	go run $(MAIN_PACKAGE)
