# Gin 安装和初始化

- 创建目录和初始化 go mod

```bash
mkdir gin_demo/chapter-01


cd gin_demo


go mod init test


cd chapter-01


touch main.go
```

- 安装 gin

```bash
go get -u github.com/gin-gonic/gin
```

- 安装 gowatch 进行热编译

```bash
go get github.com/silenceper/gowatch



# 使用



gowatch main.go
```

- Hello, Gin

```go
package main



import (

    "github.com/gin-gonic/gin"

    "net/http"

)



func main() {

    router := gin.Default()

    router.GET("/", func(c *gin.Context) {

        c.String(http.StatusOK, "Hello World")

    })

    router.Run(":8000")

}
```

运行：

```bash
gowatch main.go
```

测试：新建 `test.http` 结合 vs code 的 rest-client 插件 进行测试

```
GET http://localhost:8000/ HTTP/1.1
```
