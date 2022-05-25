postgres:
	docker run --name postgres14 --network app-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root bank_app

dropdb:
	docker exec -it postgres14 dropdb bank_app

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable" -verbose up

AWSmigrateup:
	migrate -path db/migration -database "postgresql://root:appbank.cglffkopv3rd.ap-northeast-1.rds.amazonaws.com:5432/bank_app" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/sRRRs-7/golang-postgresql/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 AWSmigrateup sqlc test server mock

