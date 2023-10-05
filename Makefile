# Nome do executável
BINARY_NAME=aws-waf-header-analyzer

build:
	go build -o $(BINARY_NAME) main.go

run:
	go run main.go

test:
	go test ./...

clean:
	go clean
	rm -f $(BINARY_NAME)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) main.go

build-freebsd:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o $(BINARY_NAME) main.go

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe main.go

build-all: build-linux build-freebsd build-windows

default: build run