.PHONY: server
server:
	go build -o build/server example/server/main.go

.PHONY: client
client:
	go build -o build/client example/client/main.go

.PHONY: clean
clean:
	go clean -i
