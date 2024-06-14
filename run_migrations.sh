#!/bin/sh
# run_migrations.sh

# Navigate to the db folder and run the migrations
cd db && sql-migrate up -config=dbconfig.yml
