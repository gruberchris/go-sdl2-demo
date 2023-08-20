build:
	go build -o bin/$(NAME) -v ./$(NAME)

run:
	go run ./$(NAME)