build:
	go build -o bin/main main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/main main.go

run:
	go run main.go