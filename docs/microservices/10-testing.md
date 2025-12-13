---
title: 微服务测试完整指南
difficulty: advanced
duration: "4-6小时"
prerequisites: ["Go 基础", "gRPC", "微服务架构"]
tags: ["测试", "单元测试", "集成测试", "Mock"]
---

# 微服务测试完整指南

本指���介绍如何为 Go 微服务编写高质量的测试，涵盖单元测试、集成测试、性能测试等各个方面。

## 📋 学习目标

完成本教程后，你将能够：

- [x] 使用 In-Memory Server 测试 gRPC 服务
- [x] 使用 httptest 测试 HTTP API
- [x] 使用 Mock 和 Stub 隔离依赖
- [x] 编写数据库测试
- [x] 进行集成测试和 E2E 测试
- [x] 进行性能测试和基准测试
- [x] 掌握测试最佳实践

## 🧪 测试金字塔

```
       /\
      /  \  E2E Tests (10%)
     /────\  慢、脆弱、昂贵
    /      \ Integration Tests (20%)
   /────────\  中等速度和成本
  /          \ Unit Tests (70%)
 /────────────\  快速、稳定、便宜
```

**原则**：
- 70% 单元测试：快速、隔离、可靠
- 20% 集成测试：测试组件交互
- 10% E2E 测试：测试完整流程

## 🔧 gRPC 服务测试

### In-Memory gRPC Server

使用 `bufconn` 创建内存中的 gRPC 服务器，无需真实网络连接：

```go
import (
    "google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
    lis = bufconn.Listen(bufSize)
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &GreeterServer{})
    go func() {
        if err := s.Serve(lis); err != nil {
            panic(err)
        }
    }()
}

func bufDialer(context.Context, string) (net.Conn, error) {
    return lis.Dial()
}
```

**优点**：
- ✅ 快速（无网络 I/O）
- ✅ 隔离（不依赖外部资源）
- ✅ 可重复（每次运行结果一致）

### 测试简单 RPC

```go
func TestSayHello(t *testing.T) {
    ctx := context.Background()

    conn, err := grpc.DialContext(ctx, "bufnet",
        grpc.WithContextDialer(bufDialer),
        grpc.WithTransportCredentials(insecure.NewCredentials()))
    require.NoError(t, err)
    defer conn.Close()

    client := pb.NewGreeterClient(conn)

    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"正常请求", "Alice", "Hello, Alice!"},
        {"空字符串", "", "Hello, !"},
        {"中文", "张三", "Hello, 张三!"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            resp, err := client.SayHello(ctx, &pb.HelloRequest{
                Name: tt.input,
            })
            assert.NoError(t, err)
            assert.Equal(t, tt.expected, resp.GetMessage())
        })
    }
}
```

### 测试流式 RPC

#### 服务端流式

```go
func TestServerStreaming(t *testing.T) {
    // ... setup ...

    stream, err := client.SayHelloStream(ctx, &pb.HelloRequest{Name: "Bob"})
    assert.NoError(t, err)

    count := 0
    for {
        resp, err := stream.Recv()
        if err == io.EOF {
            break
        }
        assert.NoError(t, err)
        count++
        assert.Contains(t, resp.GetMessage(), "Bob")
    }

    assert.Equal(t, 5, count, "应该接收到 5 条消息")
}
```

#### 客户端流式

```go
func TestClientStreaming(t *testing.T) {
    // ... setup ...

    stream, err := client.CollectHello(ctx)
    assert.NoError(t, err)

    names := []string{"Alice", "Bob", "Charlie"}
    for _, name := range names {
        err := stream.Send(&pb.HelloRequest{Name: name})
        assert.NoError(t, err)
    }

    resp, err := stream.CloseAndRecv()
    assert.NoError(t, err)
    assert.Contains(t, resp.GetMessage(), "3")
}
```

#### 双向流式

```go
func TestBidirectionalStreaming(t *testing.T) {
    // ... setup ...

    stream, err := client.ChatHello(ctx)
    assert.NoError(t, err)

    waitc := make(chan struct{})
    receivedCount := 0

    // 接收 goroutine
    go func() {
        for {
            _, err := stream.Recv()
            if err == io.EOF {
                close(waitc)
                return
            }
            assert.NoError(t, err)
            receivedCount++
        }
    }()

    // 发送消息
    names := []string{"David", "Eve", "Frank"}
    for _, name := range names {
        err := stream.Send(&pb.HelloRequest{Name: name})
        assert.NoError(t, err)
    }

    stream.CloseSend()
    <-waitc

    assert.Equal(t, len(names), receivedCount)
}
```

### 测试超时和错误

```go
func TestTimeout(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
    defer cancel()

    stream, err := client.SayHelloStream(ctx, &pb.HelloRequest{Name: "Test"})
    assert.NoError(t, err)

    count := 0
    for {
        _, err := stream.Recv()
        if err != nil {
            assert.Error(t, err)
            break
        }
        count++
    }

    assert.Less(t, count, 5, "超时应该中断流式传输")
}
```

## 🌐 HTTP API 测试

### 使用 httptest

```go
func TestHTTPHandler(t *testing.T) {
    r := gin.Default()
    r.GET("/api/hello", helloHandler)

    tests := []struct {
        name         string
        queryParam   string
        expectedCode int
        expectedBody string
    }{
        {"正常请求", "name=Alice", 200, "Alice"},
        {"无参数", "", 200, "World"},
        {"中文参数", "name=张三", 200, "张三"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req, _ := http.NewRequest("GET", "/api/hello?"+tt.queryParam, nil)
            w := httptest.NewRecorder()

            r.ServeHTTP(w, req)

            assert.Equal(t, tt.expectedCode, w.Code)
            assert.Contains(t, w.Body.String(), tt.expectedBody)
        })
    }
}
```

### 测试中间件

```go
func TestAuthMiddleware(t *testing.T) {
    r := gin.Default()
    r.Use(authMiddleware())
    r.GET("/protected", protectedHandler)

    tests := []struct {
        name         string
        token        string
        expectedCode int
    }{
        {"有效 Token", "valid-token", 200},
        {"无效 Token", "invalid-token", 401},
        {"缺少 Token", "", 401},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req, _ := http.NewRequest("GET", "/protected", nil)
            if tt.token != "" {
                req.Header.Set("Authorization", "Bearer "+tt.token)
            }

            w := httptest.NewRecorder()
            r.ServeHTTP(w, req)

            assert.Equal(t, tt.expectedCode, w.Code)
        })
    }
}
```

## 🔌 Mock 和 Stub

### 使用 testify/mock

```go
import "github.com/stretchr/testify/mock"

// 定义 Mock
type MockUserService struct {
    mock.Mock
}

func (m *MockUserService) GetUser(id int) (*User, error) {
    args := m.Called(id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*User), args.Error(1)
}

// 测试
func TestUserController(t *testing.T) {
    mockService := new(MockUserService)

    // 设置期望
    mockService.On("GetUser", 1).Return(&User{
        ID:   1,
        Name: "Test",
    }, nil)

    controller := NewUserController(mockService)
    user, err := controller.HandleGetUser(1)

    assert.NoError(t, err)
    assert.Equal(t, "Test", user.Name)

    // 验证期望被调用
    mockService.AssertExpectations(t)
}
```

### Mock HTTP Client

```go
type MockHTTPClient struct {
    DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
    return m.DoFunc(req)
}

func TestServiceCall(t *testing.T) {
    mockClient := &MockHTTPClient{
        DoFunc: func(req *http.Request) (*http.Response, error) {
            return &http.Response{
                StatusCode: 200,
                Body:       io.NopCloser(strings.NewReader(`{"status": "ok"}`)),
            }, nil
        },
    }

    service := NewService(mockClient)
    result, err := service.CallExternalAPI()

    assert.NoError(t, err)
    assert.Equal(t, "ok", result.Status)
}
```

## 💾 数据库测试

### SQLite 内存数据库

```go
func setupTestDB(t *testing.T) *sql.DB {
    db, err := sql.Open("sqlite3", ":memory:")
    require.NoError(t, err)

    _, err = db.Exec(`
        CREATE TABLE users (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL
        )
    `)
    require.NoError(t, err)

    return db
}

func TestCreateUser(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()

    userRepo := NewUserRepository(db)

    user := &User{
        Name:  "Test User",
        Email: "test@example.com",
    }

    err := userRepo.Create(user)
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)

    // 验证插入成功
    found, err := userRepo.FindByID(user.ID)
    assert.NoError(t, err)
    assert.Equal(t, user.Name, found.Name)
}
```

### 使用 testcontainers

```go
import (
    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/wait"
)

func setupPostgres(t *testing.T) (*sql.DB, func()) {
    ctx := context.Background()

    req := testcontainers.ContainerRequest{
        Image:        "postgres:14",
        ExposedPorts: []string{"5432/tcp"},
        Env: map[string]string{
            "POSTGRES_PASSWORD": "test",
            "POSTGRES_DB":       "testdb",
        },
        WaitingFor: wait.ForListeningPort("5432/tcp"),
    }

    container, err := testcontainers.GenericContainer(ctx,
        testcontainers.GenericContainerRequest{
            ContainerRequest: req,
            Started:          true,
        })
    require.NoError(t, err)

    host, _ := container.Host(ctx)
    port, _ := container.MappedPort(ctx, "5432")

    dsn := fmt.Sprintf("host=%s port=%s user=postgres password=test dbname=testdb sslmode=disable",
        host, port.Port())
    db, err := sql.Open("postgres", dsn)
    require.NoError(t, err)

    cleanup := func() {
        db.Close()
        container.Terminate(ctx)
    }

    return db, cleanup
}
```

## 🔗 集成测试

### 测试多服务交互

```go
func TestServiceIntegration(t *testing.T) {
    // 启动测试数据库
    db := setupTestDB(t)
    defer db.Close()

    // 启动 user service
    userService := NewUserService(db)
    userServer := startTestServer(userService)
    defer userServer.Close()

    // 启动 order service（依赖 user service）
    orderService := NewOrderService(userServer.URL, db)

    // 创建用户
    user := &User{Name: "Test", Email: "test@example.com"}
    err := userService.CreateUser(user)
    require.NoError(t, err)

    // 创建订单（会调用 user service 验证用户）
    order := &Order{
        UserID: user.ID,
        Amount: 100,
    }
    err = orderService.CreateOrder(order)
    assert.NoError(t, err)
}
```

## ⚡ 性能测试

### Benchmark 测试

```go
func BenchmarkSayHello(b *testing.B) {
    ctx := context.Background()
    conn, _ := grpc.DialContext(ctx, "bufnet",
        grpc.WithContextDialer(bufDialer),
        grpc.WithTransportCredentials(insecure.NewCredentials()))
    defer conn.Close()

    client := pb.NewGreeterClient(conn)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        client.SayHello(ctx, &pb.HelloRequest{Name: "Benchmark"})
    }
}

// 并行 Benchmark
func BenchmarkParallel(b *testing.B) {
    // ... setup ...

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            client.SayHello(ctx, &pb.HelloRequest{Name: "Benchmark"})
        }
    })
}
```

**运行基准测试**：

```bash
# 运行所有 benchmark
go test -bench=. -benchmem

# 运行特定 benchmark
go test -bench=BenchmarkSayHello -benchtime=10s

# 生成 CPU profile
go test -bench=. -cpuprofile=cpu.prof

# 生成内存 profile
go test -bench=. -memprofile=mem.prof
```

### 压力测试

```go
func TestHighLoad(t *testing.T) {
    const (
        concurrency = 100
        requests    = 10000
    )

    var wg sync.WaitGroup
    errors := make(chan error, requests)

    for i := 0; i < concurrency; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < requests/concurrency; j++ {
                _, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Load"})
                if err != nil {
                    errors <- err
                }
            }
        }()
    }

    wg.Wait()
    close(errors)

    errorCount := 0
    for range errors {
        errorCount++
    }

    // 允许少量错误（< 1%）
    assert.Less(t, errorCount, requests/100)
}
```

## 💡 测试最佳实践

### 1. Table-Driven Tests

```go
func TestUserValidation(t *testing.T) {
    tests := []struct {
        name    string
        user    User
        wantErr bool
    }{
        {"有效用户", User{Name: "Test", Email: "test@example.com"}, false},
        {"无效邮箱", User{Name: "Test", Email: "invalid"}, true},
        {"空名字", User{Name: "", Email: "test@example.com"}, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validateUser(tt.user)
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
```

### 2. Setup 和 Teardown

```go
func TestMain(m *testing.M) {
    // 全局 setup
    setup()

    // 运行测试
    code := m.Run()

    // 全局 teardown
    teardown()

    os.Exit(code)
}
```

### 3. 测试隔离

```go
func TestIsolation(t *testing.T) {
    // 每个测试使用独立的数据库
    db := setupTestDB(t)
    defer db.Close()

    // 测试开始前清理数据
    cleanupTestData(db)

    // 运行测试...
}
```

### 4. 使用 Fixtures

```go
// fixtures.go
func CreateTestUser(t *testing.T, db *sql.DB) *User {
    user := &User{
        Name:  "Test User",
        Email: fmt.Sprintf("test_%d@example.com", time.Now().UnixNano()),
    }
    err := db.Create(user)
    require.NoError(t, err)
    return user
}

// 使用
func TestOrderCreation(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()

    user := CreateTestUser(t, db)
    order := CreateTestOrder(t, db, user.ID)

    // 测试...
}
```

## 📊 测试覆盖率

```bash
# 生成覆盖率报告
go test -coverprofile=coverage.out ./...

# 查看覆盖率
go tool cover -func=coverage.out

# 生成 HTML 报告
go tool cover -html=coverage.out -o coverage.html
```

## 🔧 测试工具推荐

| 工具 | 用途 | 链接 |
|------|------|------|
| testify | 断言和 Mock | [github.com/stretchr/testify](https://github.com/stretchr/testify) |
| gomock | Mock 生成 | [github.com/golang/mock](https://github.com/golang/mock) |
| testcontainers | Docker 容器测试 | [github.com/testcontainers/testcontainers-go](https://github.com/testcontainers/testcontainers-go) |
| httptest | HTTP 测试 | 标准库 |
| go-sqlmock | SQL Mock | [github.com/DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) |

## 🚀 完整示例

完整的测试示例代码在：
- [gRPC 流式测试](https://github.com/Jsmond2016/go-study/blob/main/examples/microservices/01-grpc/streaming/server_test.go)
- [微服务测试指南](https://github.com/Jsmond2016/go-study/blob/main/examples/microservices/TESTING.md)

## 📚 扩展阅读

- [Go Testing 官方文档](https://pkg.go.dev/testing)
- [Testify 文档](https://pkg.go.dev/github.com/stretchr/testify)
- [Google Test Blog](https://testing.googleblog.com/)

## ⏭️ 下一步

- [gRPC 流式通信](./01-streaming) - 学习 gRPC 流式 RPC
- [服务发现](./03-service-discovery) - 服务注册与发现
- [API 网关](./05-api-gateway) - API 网关开发

---

**🎉 恭喜！** 你已经掌握了微服务测试的完整知识体系！
