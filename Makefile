build:
	@go build -o bin/friendlorant

run:
	@./bin/friendlorant/cmd/main.go

test:
	@go test -v ./...