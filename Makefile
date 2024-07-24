build:
	@go build -o bin/firstgobackendapi cmd/main.go

test:
	@go test -v ./...

run:
	@go run cmd/main.go