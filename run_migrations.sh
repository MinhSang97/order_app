#!/bin/sh

### Wait for the PostgreSQL database to be ready
#until pg_isready -h ${DB_HOST} -p ${DB_PORT} -U ${DB_USER}; do
#  echo "Waiting for PostgreSQL to be ready..."
#  sleep 5
#done

# Hiển thị thông tin thư mục hiện tại
echo "Current directory: $(pwd)"

# Điều hướng đến thư mục db và chạy các migration
cd db && sql-migrate up -config=dbconfig.yml

## Kiểm tra kết quả của lệnh cd
#if [ $? -ne 0 ]; then
#  echo "Failed to change directory to db"
#  exit 1
#fi

# Quay lại thư mục gốc
#cd ..

## Hiển thị thông tin thư mục hiện tại
#echo "Current directory after migrations: $(pwd)"



## Kiểm tra kết quả của lệnh cd
#if [ $? -ne 0 ]; then
#  echo "Failed to change directory to cmd"
#  exit 1
#fi
