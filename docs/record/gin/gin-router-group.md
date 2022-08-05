## gin-router-group

Code:

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Simple group: v1
    v1 := router.Group("/v1")
    {
        v1.GET("/login", loginEndpoint)
        v1.GET("/submit", submitEndpoint)
        v1.POST("/read", readEndpoint)
    }

    // Simple group: v2
    v2 := router.Group("/v2")
    {
        v2.POST("/login", loginEndpoint)
        v2.POST("/submit", submitEndpoint)
        v2.POST("/read", readEndpoint)
    }

    router.Run(":8080")
}
func loginEndpoint(c *gin.Context) {
    name := c.DefaultQuery("name", "Guest") //可设置默认值
    c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}

func submitEndpoint(c *gin.Context) {
    name := c.DefaultQuery("name", "Guest") //可设置默认值
    c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}

func readEndpoint(c *gin.Context) {
    name := c.DefaultQuery("name", "Guest") //可设置默认值
    c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}
```

测试：

```http
@base = 127.0.0.1:8080

### 

GET http://{{base}}/v1/login?name=hanru HTTP/1.1

###
POST http://{{base}}/form?name=测试人员&num=2 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

username=Tony
&password=123455
&hobby="basketball"
&hobby="running"
```
