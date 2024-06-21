# Use the official Golang image as a base
FROM golang:1.21-alpine

# Set the Current Working Directory inside the container
WORKDIR /order_app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

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
