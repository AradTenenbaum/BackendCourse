DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
AWS_DB_URL=postgresql://root:Femu9LcQ3szk8KVht7TD@simple-bank.cqbf8tifnnnx.eu-west-1.rds.amazonaws.com:5432/simple_bank

network:
	docker network create bank-network

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup-aws:
	migrate -path db/migration -database "$(AWS_DB_URL)" -verbose up

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

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

test:
	go test -v -cover ./...

server: 
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/AradTenenbaum/BackendCourse/db/sqlc Store

proto:
	rm -f pb/*.go 
	rm -f doc/swagger/*.swagger.json 
	rm -f doc/statik/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc

grpc-test-client:
	npm start --prefix grpc_test_client

.PHONY: network postgres createdb migrateup migrateup1 migratedown migratedown1 sqlc-init sqlc-compile sqlc-generate, db_docs, db_schema, test, server, mock, proto