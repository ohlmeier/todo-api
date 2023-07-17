.PHONY: clean
clean:
	@rm -rf bin

bin: clean
	@mkdir -p bin

.PHONY: download
download:
	go mod download -x

staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: lint
lint: bin staticcheck
	staticcheck -checks all,-ST1003 ./...

.PHONY: test
test:
	go test -v ./... -cover -json

.PHONY: build
build: download bin
	go build -o bin/todo main.go

.PHONY: build-linux
build-linux: download bin
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-w -s" -o bin ./...

.PHONY: run
run:
	go run main.go

.PHONY: help
help: ## show help
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
        awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ''