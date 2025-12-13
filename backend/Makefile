.PHONY: run build test migrate swagger

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test ./...

migrate:
	go run cmd/migrate/main.go

swagger:
	swag init -g cmd/server/main.go -o ./docs

install-deps:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@latest

