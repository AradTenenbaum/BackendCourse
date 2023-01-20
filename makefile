DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

network:
	docker network create bank-network

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

sqlc-init:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc init

sqlc-compile:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc compile

sqlc-generate:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: network postgres createdb migrateup migratedown sqlc-init sqlc-compile sqlc-generate