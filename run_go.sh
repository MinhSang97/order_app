#!/bin/sh

# Hiển thị thông tin thư mục hiện tại
echo "Current directory: $(pwd)"

# Điều hướng đến thư mục cmd và chạy ứng dụng Go
go run ./cmd/main.go