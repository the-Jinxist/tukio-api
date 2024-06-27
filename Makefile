DSN="host=localhost port=5432 user=user password=password dbname=tukio sslmode=disable timezone=UTC connect_timeout=5"
MIGRATION_DIR="./migrations"
run:
	go run main.go

up: 
	@env GOOSE_DBSTRING=${DSN} GOOSE_MIGRATION_DIR=${MIGRATION_DIR} goose postgres up

down: 
	@env GOOSE_DBSTRING=${DSN} GOOSE_MIGRATION_DIR=${MIGRATION_DIR} goose postgres down

db-status:
	@env GOOSE_MIGRATION_DIR=${MIGRATION_DIR} goose postgres ${DSN} status

