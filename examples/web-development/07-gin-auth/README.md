# Gin 认证示例

本示例展示了 Gin 框架的认证功能，包括 JWT Token 生成和验证。

## 运行示例

```bash
go mod tidy
go run main.go
```

## 测试

```bash
# 登录获取 token
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'

# 使用 token 访问受保护接口
curl -H "Authorization: Bearer YOUR_TOKEN" \
  http://localhost:8080/api/profile
```

## 知识点

### JWT Token

- JWT (JSON Web Token) 是一种无状态的认证方式
- Token 包含用户信息和过期时间
- 使用密钥签名，防止篡改

### 认证流程

1. 用户登录，验证用户名和密码
2. 生成 JWT Token
3. 客户端保存 Token
4. 后续请求携带 Token
5. 服务器验证 Token

### 中间件实现

- 从请求头获取 Token
- 验证 Token 有效性
- 将用户信息存储到 Context

## 练习

1. 实现 Token 刷新机制
2. 添加 Token 黑名单功能
3. 实现基于角色的访问控制（RBAC）

