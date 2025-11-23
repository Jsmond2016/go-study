.PHONY: proto
proto:
	@echo "Generating protobuf code..."
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--proto_path=basic \
		--proto_path=complex \
		--proto_path=service \
		--proto_path=versioning \
		-I/usr/local/include \
		basic/*.proto complex/*.proto service/*.proto versioning/*.proto
	@echo "Done!"

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	@find . -name "*.pb.go" -delete
	@echo "Done!"

.PHONY: basic
basic:
	@cd basic && go run main.go

.PHONY: complex
complex:
	@cd complex && go run main.go

.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

