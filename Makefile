server:
	go run core/main.go

build-cli:
	go build -o bin/meatball cli/main.go

migrate-create:
	migrate create -ext sql -dir db/migration -seq init_schema

createdb:
	docker exec -it meatball-db createdb --username=root --owner=root meatball

dropdb:
	docker exec -it meatball-db dropdb meatball

postgres:
	docker run --name meatball-db -p 1302:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=meatball -d postgres:16-alpine

test:
	go test -v ./...

migrateup:
	migrate -path db/migration -database "postgresql://root:meatball@localhost:1302/meatball?sslmode=disable" -verbose up

.PHONY: server build-cli postgres createdb test