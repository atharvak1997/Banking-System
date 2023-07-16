postgres:
	docker run --name bankSys -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=99220Avk97* -d postgres:12-alpine

createdb:
	docker exec -it bankSys  createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it bankSys  dropdb simple_bank

migrateup:
	migrate -path . -database "postgresql://root:99220Avk97*@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path . -database "postgresql://root:99220Avk97*@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path . -database "postgresql://root:99220Avk97*@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

migrateup1:
	migrate -path . -database "postgresql://root:99220Avk97*@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

test:
	go test -v -cover ./...

server:
	go run main.go

sqlc:
	sqlc generate

.PHONY: createdb dropdb postgres migrateup migratedown test server sqlc migrateup1 migratedown1