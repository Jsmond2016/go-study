#!/bin/bash

# 生成 Protocol Buffers 代码

echo "生成 Protocol Buffers 代码..."

# 创建输出目录
mkdir -p proto/book proto/user

# 生成 Book Service proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/book.proto

# 生成 User Service proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/user.proto

echo "Proto 代码生成完成！"
