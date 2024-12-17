# Makefile for managing Go application in Docker

# Run all necessary steps: build, migrate, and run
all: build migrate run

# Build the Go application
build:
	echo "Building Go application..."
	docker-compose exec app go build -o main .

# Run migrations
migrate:
	echo "Running migrations..."
	docker-compose exec app go run migrations/migration.go -direction up

# Start the application
run:
	echo "Starting Go application..."
	docker-compose exec -d app ./main

# Stop the application
stop:
	echo "Stopping Go application..."
	docker-compose exec app pkill -f 'main'

# Restart the application
restart: stop build run
	echo "Application restarted successfully."

# Clean up the Docker environment (optional)
clean:
	echo "Cleaning Docker environment..."
	docker-compose down -v
