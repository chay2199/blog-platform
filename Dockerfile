# Use the official Golang image to create a build artifact
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM golang:1.22

# Install sqlite3
RUN apt-get update && apt-get install -y sqlite3

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/database/migrations.sql /app/database/migrations.sql

# Command to run the executable
CMD ["sh", "-c", "sqlite3 /app/database/blog.db < /app/database/migrations.sql && ./main"]
