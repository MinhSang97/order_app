#!/bin/sh

# Wait for the PostgreSQL database to be ready
until pg_isready -h ${DB_HOST} -p ${DB_PORT} -U ${DB_USER}; do
  echo "Waiting for PostgreSQL to be ready..."
  sleep 5
done


# Navigate to the db folder and run the migrations
cd db && sql-migrate up -config=dbconfig.yml
