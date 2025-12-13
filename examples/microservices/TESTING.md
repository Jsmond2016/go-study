# 微服务单元测试指南

本指南介绍如何为微服务编写高质量的单元测试，包括 gRPC 服务、HTTP API、数据库操作等常见场景。

## 目录

1. [gRPC 服务测试](#grpc-服务测试)
2. [HTTP API 测试](#http-api-测试)
3. [服务间调用测试](#服务间调用测试)
4. [数据库测试](#数据库测试)
5. [Mock 和 Stub](#mock-和-stub)
6. [集成测试](#集成测试)
7. [性能测试](#性能测试)
8. [最佳实践](#最佳实践)

## gRPC 服务测试

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

func TestSayHello(t *testing.T) {
    ctx := context.Background()
    conn, _ := grpc.DialContext(ctx, "bufnet",
        grpc.WithContextDialer(bufDialer),
        grpc.WithTransportCredentials(insecure.NewCredentials()))
    defer conn.Close()

    client := pb.NewGreeterClient(conn)
    resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Test"})

    assert.NoError(t, err)
    assert.Equal(t, "Hello, Test!", resp.GetMessage())
}
```

**优点**:
- 快速（无网络 I/O）
- 隔离（不依赖外部资源）
- 可重复（每次运行结果一致）

### 流式 RPC 测试

#### 服务端流式测试

```go
func TestServerStreaming(t *testing.T) {
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

    assert.Equal(t, 5, count)
}
```

#### 客户端流式测试

```go
func TestClientStreaming(t *testing.T) {
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

#### 双向流式测试

```go
func TestBidirectionalStreaming(t *testing.T) {
    stream, err := client.ChatHello(ctx)
    assert.NoError(t, err)

    // 使用 goroutine 并发收发
    waitc := make(chan struct{})

    go func() {
        for {
            resp, err := stream.Recv()
            if err == io.EOF {
                close(waitc)
                return
            }
            assert.NoError(t, err)
            // 验证响应...
        }
    }()

    // 发送请求
    for _, name := range names {
        err := stream.Send(&pb.HelloRequest{Name: name})
        assert.NoError(t, err)
    }

    stream.CloseSend()
    <-waitc
}
```

### 超时和错误处理测试

```go
func TestTimeout(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
    defer cancel()

    stream, err := client.SayHelloStream(ctx, &pb.HelloRequest{Name: "Test"})
    assert.NoError(t, err)

    // 应该在超时前只收到部分消息
    count := 0
    for {
        _, err := stream.Recv()
        if err != nil {
            assert.Error(t, err)
            break
        }
        count++
    }

    assert.Less(t, count, 5)
}
```

### 并发测试

```go
func TestConcurrentCalls(t *testing.T) {
    const concurrency = 100
    done := make(chan bool, concurrency)

    for i := 0; i < concurrency; i++ {
        go func(id int) {
            resp, err := client.SayHello(ctx, &pb.HelloRequest{
                Name: fmt.Sprintf("User-%d", id),
            })
            assert.NoError(t, err)
            assert.Contains(t, resp.GetMessage(), fmt.Sprintf("User-%d", id))
            done <- true
        }(i)
    }

    for i := 0; i < concurrency; i++ {
        <-done
    }
}
```

## HTTP API 测试

### 使用 httptest

```go
func TestHTTPHandler(t *testing.T) {
    r := gin.Default()
    r.GET("/api/hello", helloHandler)

    // 创建测试请求
    req, _ := http.NewRequest("GET", "/api/hello?name=Test", nil)
    w := httptest.NewRecorder()

    // 执行请求
    r.ServeHTTP(w, req)

    // 验证响应
    assert.Equal(t, http.StatusOK, w.Code)

    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "Hello, Test!", response["message"])
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

## 服务间调用测试

### Mock HTTP Client

```go
// 定义 HTTP Client 接口
type HTTPClient interface {
    Do(req *http.Request) (*http.Response, error)
}

// Mock 实现
type MockHTTPClient struct {
    DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
    return m.DoFunc(req)
}

// 测试
func TestServiceCall(t *testing.T) {
    mockClient := &MockHTTPClient{
        DoFunc: func(req *http.Request) (*http.Response, error) {
            // 返回模拟响应
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

### 使用 httptest.Server

```go
func TestExternalServiceCall(t *testing.T) {
    // 创建测试服务器
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "success"}`))
    }))
    defer server.Close()

    // 使用测试服务器的 URL
    client := NewClient(server.URL)
    resp, err := client.GetData()

    assert.NoError(t, err)
    assert.Equal(t, "success", resp.Message)
}
```

## 数据库测试

### 使用 SQLite 内存数据库

```go
func setupTestDB(t *testing.T) *sql.DB {
    db, err := sql.Open("sqlite3", ":memory:")
    require.NoError(t, err)

    // 创建表
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

### 使用 testcontainers（真实数据库）

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

    container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
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

## Mock 和 Stub

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

### 使用 GoMock

```go
// 生成 mock
//go:generate mockgen -destination=mocks/user_service.go -package=mocks . UserService

type UserService interface {
    GetUser(ctx context.Context, id int) (*User, error)
}

// 测试
func TestWithGoMock(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockService := mocks.NewMockUserService(ctrl)

    mockService.EXPECT().
        GetUser(gomock.Any(), 1).
        Return(&User{ID: 1, Name: "Test"}, nil)

    // 使用 mock...
}
```

## 集成测试

### 测试多个服务交互

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

## 性能测试

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

func BenchmarkParallel(b *testing.B) {
    // ... setup ...

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            client.SayHello(ctx, &pb.HelloRequest{Name: "Benchmark"})
        }
    })
}
```

运行基准测试：

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

    // 允许少量错误（如< 1%）
    assert.Less(t, errorCount, requests/100)
}
```

## 最佳实践

### 1. 测试组织

```go
// 使用 table-driven tests
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

func setup() {
    // 初始化测试环境
}

func teardown() {
    // 清理测试环境
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

### 4. 使用测试 Fixtures

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

### 5. 测试覆盖率

```bash
# 生成覆盖率报告
go test -coverprofile=coverage.out ./...

# 查看覆盖率
go tool cover -func=coverage.out

# 生成 HTML 报告
go tool cover -html=coverage.out -o coverage.html
```

### 6. 持续集成

```yaml
# .github/workflows/test.yml
name: Tests
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'

      - name: Run tests
        run: go test -v -race -coverprofile=coverage.out ./...

      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          file: ./coverage.out
```

## 测试工具推荐

1. **testify**: 断言和 mock 库
2. **gomock**: Google 的 mock 框架
3. **testcontainers**: Docker 容器测试
4. **httptest**: HTTP 测试工具
5. **go-sqlmock**: SQL mock 库
6. **gock**: HTTP mock 库

## 总结

编写好的微服务测试需要：

- ✅ 测试隔离（独立数据库、内存服务器）
- ✅ 快速反馈（使用 mock 和 stub）
- ✅ 完整覆盖（单元、集成、E2E）
- ✅ 性能测试（benchmark 和压力测试）
- ✅ 持续集成（自动化测试流程）

遵循这些最佳实践，可以构建健壮、可靠的微服务系统。
