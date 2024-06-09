.PHONY: migrate-up migrate-down

# Docker commands:
docker up:
	docker compose up -d

docker down:
	docker compose down

migrate-up:
	cd db && sql-migrate up -config=dbconfig.yml

migrate-down:
	cd db && sql-migrate down -config=dbconfig.yml

migrate-status:
	cd db && sql-migrate status -config=dbconfig.yml

migrate-new:
	@read -p "Enter migration name: " name; \
	timestamp=$$(date +"%Y%m%d%H%M%S"); \
	echo "-- +migrate Up\n\n\n-- +migrate Down\n" > db/migrations/$$timestamp-$$name.sql; \
	echo "Created migration file: db/migrations/$$timestamp-$$name.sql"

#Run server:
run:
	cd cmd && go run main.go