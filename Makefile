# Simple Makefile for a Go project

include .env
export $(shell sed 's/=.*//' .env)

# Build the application
all: build

build:
	@echo "Building..."
	
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go


# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi


# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v


# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v


# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

tools:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

generate:
	sqlc generate

# Migrate the database
db-migrate: tools
	@echo "Migrating database..."
	@migrate -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path ./internal/database/migrations up

# Create a new migration
migration-add:
	@if [ -z "$(name)" ]; then \
		echo "Please provide a migration name. Usage: make migration-add name=<migration_name>"; \
		exit 1; \
	fi
	@migrate create -ext sql -path internal/database/migrations -seq $(name)

.PHONY: all build run test clean db-migrate
