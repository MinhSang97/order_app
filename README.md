```markdown
# Order App

Order App is an application that allows users to place orders online.

## Quick Start

- Install [Go](https://golang.org/dl/)
- Install [Docker](https://docs.docker.com/get-docker/)
- SET environment variables in .env file
- SET environment Go

## Clone project

```bash
git clone https://github.com/MinhSang97/order_app.git
```

## Move to oneship folder

```bash
cd order_app
```

## Install dependencies

```bash
go mod tidy
```

## DB Migration

```bash
make migrate-up
```

## Run project

```bash
go run ./cmd/main.go
```

## LICENSE

This project is distributed under the MIT License. See the `LICENSE` file for more information.
```
