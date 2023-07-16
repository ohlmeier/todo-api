build:
	go build -o bin/todo main.go

build-linux:
	set GOOS=linux
	set GOARCH=amd64
	go build -o bin/todo main.go

run:
	go run main.go