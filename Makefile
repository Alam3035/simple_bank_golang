postgres:
	docker-compose up -d --build postgres

stop:
	docker-compose down -v

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

start:
	docker-compose up -d --build postgres && sleep 15 && docker exec -it postgres12 createdb --username=root --owner=root simple_bank && migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up && go run main.go

.PHONY: postgres stopgres createdb dropdb migrateup migratedown sqlc server