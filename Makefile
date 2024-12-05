include .env

MIGRATION_PATH=./cmd/migrate/migrations


migration:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))


migrate-up:
	@migrate -path $(MIGRATION_PATH) -database "$(DB_ADDR)" up


migrate-down:
	@migrate -path $(MIGRATION_PATH) -database "$(DB_ADDR)" down $(filter-out $@,$(MAKECMDGOALS))


migrate-reset:
	migrate -path $(MIGRATION_PATH) -database "$(DB_ADDR)" down -all

# Drop migration tables and recreate
migrate-fresh:
	psql "$(DB_ADDR)" -c "DROP TABLE IF EXISTS schema_migrations, users;"
	migrate -path $(MIGRATION_PATH) -database "$(DB_ADDR)" up

.PHONY: migrate-create migrate-up migrate-down migrate-reset migrate-fresh