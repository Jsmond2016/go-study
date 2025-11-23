.PHONY: proto clean

proto:
	@echo "生成 Go 代码..."
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		user.proto order.proto product.proto

clean:
	@echo "清理生成的文件..."
	rm -rf user/*.pb.go
	rm -rf order/*.pb.go
	rm -rf product/*.pb.go

