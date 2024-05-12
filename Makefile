build:
	@go build -o bin/friendlorant

run:
	@./bin/friendlorant

test:
	@go test -v ./...