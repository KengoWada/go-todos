include env/app.env
MIGRATIONS_PATH = ./cmd/migrate/migrations

migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) up

migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database $(DB_ADDR) down

test-migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database $(TEST_DB_ADDR) up

test-migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database $(TEST_DB_ADDR) down

services-up:
	@docker compose -f docker-compose.yml up --build

services-down:
	@docker compose -f docker-compose.yml down

services-kill:
	@docker compose -f docker-compose.yml down -v

.PHONY: migration services-up services-down services-kill migrate-up migrate-down test-migrate-up test-migrate-down