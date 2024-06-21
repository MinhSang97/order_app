## Use the official Golang image as a base
#FROM golang:1.21-alpine
#
## Install the curl package
#RUN apk add --no-cache curl
#
## Install the postgresql-client package
#RUN curl -OL https://golang.org/dl/go1.22.3.linux-amd64.tar.gz && \
#    tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz && \
#    rm go1.22.3.linux-amd64.tar.gz && \
#    export PATH=$PATH:/usr/local/go/bin
#
## Set the Current Working Directory inside the container
#WORKDIR /order_app
#
## Copy go mod and sum files
#COPY go.mod go.sum ./
#
## Install dependencies
#RUN /usr/local/go/bin/go mod tidy
#
## Install sql-migrate
#RUN go install github.com/rubenv/sql-migrate/...@latest
#
## Copy the source code into the container
#COPY . .
#
## Set environment variables
#ENV GO_ENV=development
#
## Expose port 8080 to the outside world
#EXPOSE 8080
#
## Run the migrations
#RUN chmod +x ./run_migrations.sh
#RUN ./run_migrations.sh
#
## Command to run the application
#CMD ["go", "run", "./cmd/main.go"]


## Use a specific version of Go that matches your requirements
#FROM golang:1.22-alpine
#
## Install curl and git
#RUN apk add --no-cache curl git
#
## Set Go environment variable
#ENV PATH /usr/local/go/bin:$PATH
#
#WORKDIR /order_app
#
## Copy go mod files and install dependencies
#COPY go.mod go.sum ./
#RUN go mod tidy
#
## Install sql-migrate
#RUN go install github.com/rubenv/sql-migrate/...@latest
#
## Copy the rest of the application files
#COPY . .
#
## Ensure the migration script is executable
#RUN chmod +x ./run_migrations.sh
#
## Set environment variables
#ENV GO_ENV=development
#
## Expose the application port
#EXPOSE 8080
#
## Run the migration script
#RUN ./run_migrations.sh

# Use a specific version of Go that matches your requirements
FROM golang:1.22-alpine

# Install curl, git, and PostgreSQL client tools
RUN apk add --no-cache curl git postgresql-client

# Set Go environment variable
ENV PATH /usr/local/go/bin:$PATH

WORKDIR /order_app

# Copy go mod files and install dependencies
COPY go.mod go.sum ./

RUN go mod tidy

# Install sql-migrate
RUN go install github.com/rubenv/sql-migrate/...@latest

RUN docker compose up -d

# Copy the rest of the application files
COPY . .

# Ensure the migration script is executable
RUN chmod +x ./run_migrations.sh

# Set environment variables
ENV GO_ENV=development

# Expose the application port
EXPOSE 8080

# Run the migration script
CMD ["./run_migrations.sh"]


