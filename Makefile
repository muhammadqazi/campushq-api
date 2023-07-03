DB_URL=$(shell grep DB_URL .env | cut -d '=' -f2-)

postgres:
	docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:latest
create-db:
	docker exec -it postgres createdb --username=root --owner=root campushq
db-wipe:
	docker exec -it postgres dropdb campushq
db-init:
	echo $(DB_URL)
	@cd src/internal/core/infrastructure/postgres && \
	docker exec -it postgres createdb --username=root --owner=root campushq && \
	docker exec -i postgres psql -U root -d campushq < database-up.sql && \
	docker exec -i postgres psql -U root -d campushq < db-seed.sql
db-seed:
	@cd src/internal/core/infrastructure/postgres && \
	docker exec -i postgres psql -U root -d campushq < db-seed.sql
new-migration:
	migrate create -ext sql -dir  src/internal/core/infrastructure/postgres/migrations -seq $(name)
migrate-up:
	migrate -path src/internal/core/infrastructure/postgres/migrations -database "$(DB_URL)" -verbose up
migrate-down:
	migrate -path src/internal/core/infrastructure/postgres/migrations -database "$(DB_URL)" -verbose down
sqlc:
	sqlc generate
server:
	@if [ -z $$(docker exec -it postgres psql -U root -lqt | cut -d \| -f 1 | grep -w "campushq") ]; then \
		make create-db; \
	fi
	go run src/cmd/main.go

.PHONY: postgres create-db db-wipe db-init db-seed server new-migration migrate-up migrate-down sqlc