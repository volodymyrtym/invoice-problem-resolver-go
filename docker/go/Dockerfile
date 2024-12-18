# --- Stage 1: Builder (development and application build) ---
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy Go modules to cache dependencies
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download

# Copy the rest of the sgo.sum ource files
COPY . .

# For production, compile the final application
RUN go build -o main

# --- Stage 2: Development ---
FROM golang:1.23 AS dev

# Install Air for live reload during development
RUN go install github.com/cosmtrek/air@v1.61.2

# Set the working directory
WORKDIR /app

# Copy the entire local codebase (used in development mode)
COPY . .

# Expose the port for the application
EXPOSE 8080

# Start the application using Air for live reload
CMD ["air"]

# --- Stage 3: Production ---
FROM gcr.io/distroless/base AS prod

# Set the working directory
WORKDIR /app

# Copy the compiled application from the builder stage
COPY --from=builder /app/main .

# Expose the port for the production application
EXPOSE 8080

# Command to start the application in production mode
CMD ["./app"]