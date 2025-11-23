.PHONY: proto
proto:
	@echo "Generating protobuf code..."
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--proto_path=proto \
		proto/*.proto
	@echo "Done!"

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	@rm -f proto/*.pb.go
	@echo "Done!"

.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

.PHONY: server
server:
	@go run server/main.go

.PHONY: client
client:
	@go run client/main.go

