build:
	@go build -o bin/api

run: build
	@./bin/api

dev: 
	@nodemon --exec go run main.go --signal SIGTERM