# Use the official Golang image as a base
FROM golang:1.21-alpine

# Install the curl package
RUN apk add --no-cache curl

# Install the postgresql-client package
RUN curl -OL https://golang.org/dl/go1.22.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz && \
    rm go1.22.3.linux-amd64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin

# Set the Current Working Directory inside the container
WORKDIR /order_app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Install dependencies
RUN /usr/local/go/bin/go mod tidy

# Install sql-migrate
RUN go install github.com/rubenv/sql-migrate/...@latest

# Copy the source code into the container
COPY . .

# Set environment variables
ENV GO_ENV=development

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the migrations
RUN chmod +x ./run_migrations.sh
RUN ./run_migrations.sh

# Command to run the application
CMD ["go", "run", "./cmd/main.go"]
