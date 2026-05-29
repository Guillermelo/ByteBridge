.DEFAULT_GOAL := run

fmt:
				 go fmt ./...
vet: fmt
				go vet ./...
run: vet
				go run ./cmd/server -port=4000 -env="development"
build: vet
				go build -o bin/server ./cmd/server

runclient: vet
				go run ./cmd/client -port=4000 -env="development"
buildclient: vet
				go build -o bin/client ./cmd/client

