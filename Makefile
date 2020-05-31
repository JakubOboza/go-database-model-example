build:
	go build -o bin/example ./cmd/example
	
test:
	go test -v ./...
