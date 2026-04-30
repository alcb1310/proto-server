build:
	go build -o ./bin/proto-server ./cmd/api/main.go

run: build
	./bin/proto-server

watch:
	air

build_proto:
	protoc --go_out=. --go_opt=paths=source_relative cmd/schemas/*.proto
