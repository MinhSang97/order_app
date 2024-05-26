# .PHONY: migrate-up migrate-down

# # Pattern rule for migration
# migrate-%:
# 	$(eval CMD := $*)
# 	cd db; \
# 	sql-migrate $(CMD) -config=dbconfig.yml

# # Explicit targets for up and down migrations
# migrate-up: migrate-up
# migrate-down: migrate-down

.PHONY: migrate-up migrate-down

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