BINARY_NAME=aws-waf-header-analyzer
INSTALL_PATH=/usr/local/bin

build:
	go build -o ./bin/$(BINARY_NAME) main.go

run:
	go run main.go

test:
	go test ./...

clean:
	go clean
	rm -f $(BINARY_NAME)

install:
	install -Dm755 ./bin/$(BINARY_NAME) ${INSTALL_PATH}

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/$(BINARY_NAME) main.go

build-freebsd:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o ./bin/$(BINARY_NAME) main.go

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/$(BINARY_NAME).exe main.go

build-all: build-linux build-freebsd build-windows

default: build run
