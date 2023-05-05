DB_URL=postgresql://root:root@localhost:5432/simple_bank?sslmode=disable

network:
	docker network create simplebank-network

postgres:
	docker run --name postgres12 --network simple_bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down


migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen --package mockdb -destination db/mock/store.go github.com/grayjunzi/backend-master-class-golang/db/sqlc Store

.PHONEY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock 