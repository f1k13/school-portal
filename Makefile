db-push:
	$(MAKE) mig-down
	$(MAKE) mig-up
	$(MAKE) db-gen

# Start the application
start:
	go run ./cmd/server/main.go 

# Create a new migration
mig-new:
	PowerShell -Command "$env:GOOSE_DRIVER='$(GOOSE_DRIVER)'; $env:GOOSE_MIGRATION_DIR='$(MIGRATION_DIR)'; $env:GOOSE_DBSTRING='$(DATABASE_URL)'; goose create new-migration sql"


# Apply all migrations
mig-up:
	PowerShell -Command "$env:GOOSE_DRIVER='$(GOOSE_DRIVER)'; $env:GOOSE_MIGRATION_DIR='$(MIGRATION_DIR)'; $env:GOOSE_DBSTRING='$(DATABASE_URL)'; goose up"


# Roll back the last migration
mig-down:
	PowerShell -Command "$env:GOOSE_DRIVER='$(GOOSE_DRIVER)'; $env:GOOSE_MIGRATION_DIR='$(MIGRATION_DIR)'; $env:GOOSE_DBSTRING='$(DATABASE_URL)'; goose down"

# Generate Jet ORM code
db-gen:
	jet -dsn="postgres://postgres:1234@localhost:5432/school-portal?sslmode=disable" -schema=public -path=./internal/storage/postgres
	sh ./model_postgen.sh