# 电商系统示例

这是一个完整的电商系统示例，包含商品管理、购物车、订单系统等核心功能。

## 功能特性

- ✅ 商品管理（CRUD）
- ✅ 购物车管理
- ✅ 订单创建和管理
- ✅ 库存管理
- ✅ 订单状态流转

## 运行示例

```bash
# 安装依赖
go mod tidy

# 运行程序
go run main.go
```

## 测试 API

### 1. 创建商品

```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name":"Go 语言教程",
    "description":"一本优秀的 Go 语言教程",
    "price":99.99,
    "stock":100
  }'
```

### 2. 获取商品列表

```bash
curl http://localhost:8080/api/products
```

### 3. 添加到购物车

```bash
curl -X POST http://localhost:8080/api/cart/items \
  -H "Content-Type: application/json" \
  -d '{
    "product_id":1,
    "quantity":2
  }'
```

### 4. 获取购物车

```bash
curl http://localhost:8080/api/cart
```

### 5. 创建订单

```bash
curl -X POST http://localhost:8080/api/orders
```

### 6. 获取订单列表

```bash
curl http://localhost:8080/api/orders
```

### 7. 更新订单状态

```bash
curl -X PATCH http://localhost:8080/api/orders/1/status \
  -H "Content-Type: application/json" \
  -d '{"status":"paid"}'
```

## 项目结构

```
03-e-commerce/
├── main.go          # 主程序
├── go.mod           # 依赖管理
└── README.md        # 说明文档
```

## 扩展功能

可以在此基础上添加：
- 支付集成
- 用户认证
- 地址管理
- 优惠券系统
- 商品评价

