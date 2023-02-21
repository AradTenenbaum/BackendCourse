DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

network:
	docker network create bank-network

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup-aws:
	migrate -path db/migration -database "postgresql://root:Femu9LcQ3szk8KVht7TD@simple-bank.cqbf8tifnnnx.eu-west-1.rds.amazonaws.com:5432/simple_bank" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

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

server: 
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/AradTenenbaum/BackendCourse/db/sqlc Store

.PHONY: network postgres createdb migrateup migrateup1 migratedown migratedown1 sqlc-init sqlc-compile sqlc-generate server mock