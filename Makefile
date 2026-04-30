build:
	go build -o ./bin/proto-server ./cmd/api/main.go

run: build
	./bin/proto-server
