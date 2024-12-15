#!/bin/bash

set -e

docker compose stop

echo "Starting Docker containers..."
docker compose up -d

echo "Tidying Go dependencies..."
docker-compose exec app go mod tidy

echo "Building Go application..."
docker-compose exec app go run migration.go -direction up
docker-compose exec app go build -o main .
