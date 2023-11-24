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

.PHONY: server build-cli postgres createdb