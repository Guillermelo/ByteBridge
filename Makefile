.DEFAULT_GOAL := server

fmt:
				 go fmt ./...
vet: fmt
				go vet ./...
server: vet
				go run ./cmd/server -port=4000 -env="development"
buildserver: vet
				go build -o bin/server ./cmd/server

client: vet
				go run ./cmd/client -file=$(file)

# 				go run ./cmd/client -port=4000 -file=$(file) -env="development"
buildclient: vet
				go build -o bin/client ./cmd/client 

