DB_DSN ?= postgres://agenticitems:agenticitems@localhost:5432/agenticitems?sslmode=disable

## run/api: run the API server locally
run/api/local:
	go run ./cmd/api -db-dsn=${DB_DSN}

## run entire application stack in docker
run/app/up:
	docker compose up -d --build

## tear down the entire application stack in docker
run/app/down:
	docker compose down

## db/migrations/up: apply all up migrations
db/migrations/up:
	migrate -path=./migrations -database=${DB_DSN} up

## db/migrations/down: apply all down migrations
db/migrations/down:
	migrate -path=./migrations -database=${DB_DSN} down

## db/migrations/reset: wipe migration state and re-run all up migrations
db/migrations/reset:
	@docker compose -f ../docker-compose.yml exec db psql -U agenticitems -d agenticitems -c "DELETE FROM schema_migrations;"
	migrate -path=./migrations -database=${DB_DSN} up

## db/wipe: drop all tables and re-run all migrations from scratch
db/wipe:
	@docker compose -f ../docker-compose.yml exec db psql -U agenticitems -d agenticitems -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
	migrate -path=./migrations -database=${DB_DSN} up

## db/psql: connect to the database with psql
db/psql:
	docker compose -f ../docker-compose.yml exec db psql -U agenticitems -d agenticitems
