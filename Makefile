BINARY_NAME=go-sdl2-demo

build:
	go build -o ./bin/$(BINARY_NAME) main.go

run:
	go run main.go

clean:
	go clean
	rm -f ./bin/${BINARY_NAME}

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all