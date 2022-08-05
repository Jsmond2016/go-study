## Gin-ping 测试

```go
package main

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World")
    })
    // r.Run() // 监听并在 0.0.0.0:8080 上启动服务
    // http.ListenAndServe(":8888", router)

    // 自定义HTTP服务器的配置
    s := &http.Server{
        Addr:           ":8080",
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
```

测试代码：

```http
@base = 127.0.0.1:8080

### 

GET http://{{base}}/v1/ping HTTP/1.1
```
