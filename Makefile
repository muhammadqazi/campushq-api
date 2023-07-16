DB_URL=$(shell grep DB_URL .env | cut -d '=' -f2-)

postgres:
	docker run --name campushq-container -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:latest
create-db:
	docker exec -it campushq-container createdb --username=root --owner=root campushq
db-wipe:
	@cd src/internal/core/infrastructure/postgres && \
	docker compose down && \
	docker-compose rm -fsv && \
	docker volume rm infrastructure_vol-campushq-db
db-start:
	@cd src/internal/core/infrastructure/postgres && \
	docker compose up -d
db-init:
	@echo "Migrating database schema and seeding data..."
	make migrate-up
	@cd src/internal/core/infrastructure/postgres; docker exec -i campushq-container psql -U root -d campushq < db-seed.sql
new-migration:
	migrate create -ext sql -dir  src/internal/core/infrastructure/postgres/migrations -seq $(name)
migrate-up:
	migrate -path src/internal/core/infrastructure/postgres/migrations -database "$(DB_URL)" -verbose up
migrate-down:
	migrate -path src/internal/core/infrastructure/postgres/migrations -database "$(DB_URL)" -verbose down
sqlc:
	sqlc generate
server:
	@if [ -z $$(docker exec -it campushq-container psql -U root -lqt | cut -d \| -f 1 | grep -w "campushq") ]; then \
		make create-db; \
	fi
	go run src/cmd/main.go

.PHONY: postgres create-db db-wipe db-start db-init db-seed server new-migration migrate-up migrate-down sqlc