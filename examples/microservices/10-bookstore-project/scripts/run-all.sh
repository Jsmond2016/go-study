#!/bin/bash

# 启动所有微服务

echo "启动在线书店微服务..."

# 启动 Book Service
echo "启动 Book Service (端口 50051)..."
cd book-service && go run main.go &
BOOK_PID=$!

sleep 2

# 启动 User Service
echo "启动 User Service (端口 50052)..."
cd ../user-service && go run main.go -port 50052 &
USER_PID=$!

sleep 2

# 启动 API Gateway
echo "启动 API Gateway (端口 8080)..."
cd ../api-gateway && go run main.go &
GATEWAY_PID=$!

echo ""
echo "所有服务已启动！"
echo "Book Service PID: $BOOK_PID"
echo "User Service PID: $USER_PID"
echo "API Gateway PID: $GATEWAY_PID"
echo ""
echo "访问 http://localhost:8080 使用 API"
echo ""
echo "按 Ctrl+C 停止所有服务"

# 等待中断信号
trap "kill $BOOK_PID $USER_PID $GATEWAY_PID; exit" INT TERM

wait
