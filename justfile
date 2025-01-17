# Load .env variables
set dotenv-load := true

# Build the project
build:
	go build -o bin/minecv cmd/main.go

# Run the project
run:
	go run cmd/main.go

# Format code
fmt:
	go fmt ./...

# Run tests
test:
	go test ./... -v

# Initialize the database
db-init:
	@echo "Initializing database..."
	@psql -U $POSTGRES_DB -d postgres -c "CREATE DATABASE $POSTGRES_DB;" || echo "Database $POSTGRES_DB already exists"
	@psql -U $POSTGRES_USER -d postgres -c "GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;"

# Reset the database (drops and recreates)
db-reset:
	@echo "Resetting database..."
	@psql -U $POSTGRES_USER -d postgres -c "DROP DATABASE IF EXISTS $POSTGRES_USER;"
	@just db-init

# Run database migrations
migrate:
	migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable" -path ./internal/infrastructure/database/migrate up

# Rollback last migration
migrate-down:
	migrate -database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" -path ./migrations down 1

# Generate a new migration file
migrate-new NAME:
	migrate create -ext sql -dir ./migrations {{NAME}}

# Clean build artifacts
clean:
	rm -rf bin

# Start the application with live reload (requires `air`)
dev:
	air

# Help menu
help:
	@echo "Available commands:"
	@just --list
