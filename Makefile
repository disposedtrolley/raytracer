ENTRYPOINT:=cmd/raytracer/main.go
BIN_NAME:=ray

build:
	go build -o $(BIN_NAME) $(ENTRYPOINT)

test:
	go test -race -cover ./...
