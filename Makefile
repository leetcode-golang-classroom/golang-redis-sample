.PHONY=build
include .env
export $(shell sed 's/=.*//' .env)
build:
	@CGO_ENABLED=0 GOOS=linux go build -o bin/main cmd/main.go

run: build
	@./bin/main


coverage:
	@go test -v -cover ./...

test:
	@go test -v ./...

redis-compose:
	@sudo docker compose up -d