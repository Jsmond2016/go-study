# API 文档

## 基础信息

- **Base URL**: `http://localhost:8080/api`
- **Content-Type**: `application/json`

## 用户服务 API

### 创建用户

```http
POST /api/users
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123",
  "full_name": "Test User"
}
```

**响应**:
```json
{
  "user_id": 1,
  "username": "testuser",
  "email": "test@example.com",
  "full_name": "Test User",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

### 获取用户

```http
GET /api/users/{id}
```

### 更新用户

```http
PUT /api/users/{id}
Content-Type: application/json

{
  "email": "newemail@example.com",
  "full_name": "New Name"
}
```

### 删除用户

```http
DELETE /api/users/{id}
```

### 用户登录

```http
POST /api/users/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

**响应**:
```json
{
  "success": true,
  "token": "token-1-1234567890",
  "user": {
    "user_id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "full_name": "Test User"
  }
}
```

## 订单服务 API

### 创建订单

```http
POST /api/orders
Content-Type: application/json

{
  "user_id": 1,
  "items": [
    {
      "product_id": 1,
      "quantity": 2,
      "price": 99.99
    }
  ],
  "shipping_address": "123 Main St"
}
```

### 获取订单

```http
GET /api/orders/{id}
```

### 获取用户订单列表

```http
GET /api/orders/user/{user_id}?page=1&page_size=10
```

### 更新订单状态

```http
PUT /api/orders/{id}/status
Content-Type: application/json

{
  "status": "shipped"
}
```

### 取消订单

```http
DELETE /api/orders/{id}
```

## 商品服务 API

### 创建商品

```http
POST /api/products
Content-Type: application/json

{
  "name": "Test Product",
  "description": "Test Description",
  "price": 99.99,
  "stock": 100,
  "category": "electronics"
}
```

### 获取商品

```http
GET /api/products/{id}
```

### 更新商品

```http
PUT /api/products/{id}
Content-Type: application/json

{
  "name": "Updated Product",
  "price": 89.99,
  "stock": 90
}
```

### 删除商品

```http
DELETE /api/products/{id}
```

### 获取商品列表

```http
GET /api/products?page=1&page_size=10&category=electronics
```

## 错误响应

所有 API 在出错时返回以下格式：

```json
{
  "error": "错误描述"
}
```

## 状态码

- `200 OK`: 请求成功
- `400 Bad Request`: 请求参数错误
- `404 Not Found`: 资源不存在
- `500 Internal Server Error`: 服务器内部错误

