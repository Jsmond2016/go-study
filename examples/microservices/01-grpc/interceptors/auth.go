package interceptors

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor 认证拦截器
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 从 context 中获取 metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	// 获取 token
	tokens := md.Get("authorization")
	if len(tokens) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing authorization token")
	}

	token := tokens[0]
	
	// 简单的 token 验证（实际应用中应该使用 JWT 等）
	if token != "valid-token-123" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	// 将用户信息存储到 context
	ctx = context.WithValue(ctx, "user_id", "user-123")
	ctx = context.WithValue(ctx, "token", token)

	return handler(ctx, req)
}

// GetUserID 从 context 获取用户 ID
func GetUserID(ctx context.Context) (string, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return "", fmt.Errorf("user_id not found in context")
	}
	return userID, nil
}

