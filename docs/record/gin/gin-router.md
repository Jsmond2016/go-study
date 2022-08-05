## gin-路由

代码：

```go
package main

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/get", func(c *gin.Context) {
        c.String(http.StatusOK, "GET METHOD")
    })

    router.GET("/get-params/:name", func(c *gin.Context) {
        username := c.Param("name")
        c.String(http.StatusOK, username)
    })

    router.POST("/post", func(c *gin.Context) {
        c.String(http.StatusOK, "POST METHOD")
    })
    router.PUT("/put", func(c *gin.Context) {
        c.String(http.StatusOK, "PUT METHOD")
    })
    router.DELETE("/delete", func(c *gin.Context) {
        c.String(http.StatusOK, "DELETE METHOD")
    })
    router.PATCH("/patch", func(c *gin.Context) {
        c.String(http.StatusOK, "PATCH METHOD")
    })
    router.HEAD("/head", func(c *gin.Context) {
        c.String(http.StatusOK, "HEAD METHOD")
    })
    router.OPTIONS("/options", func(c *gin.Context) {
        c.String(http.StatusOK, "OPTIONS METHOD")
    })

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

测试：

```http
@base = 127.0.0.1:8080

### 

GET http://{{base}}/get?id=12312323&num=2 HTTP/1.1

###
GET http://{{base}}/get-params/xiaoming HTTP/1.1

###
POST http://{{base}}/post?id=12312323&num=2 HTTP/1.1
Content-Type: application/json

{
  "name": "Tony",
  "password": "123456"
}

###
PUT http://{{base}}/put?id=12312323&num=2 HTTP/1.1
Content-Type: application/json

{
  "name": "Tony",
  "password": "123456"
}

###
DELETE  http://{{base}}/delete?id=12312323&num=2 HTTP/1.1
Content-Type: application/json

{
  "id": 123
}

###
PATCH http://{{base}}/patch?id=12312323&num=2 HTTP/1.1
Content-Type: application/json

{
  "id": 123
}
###
HEAD  http://{{base}}/head?id=12312323&num=2 HTTP/1.1


###

OPTIONS http://{{base}}/options?id=12312323&num=2 HTTP/1.1
Content-Type: application/json

{
  "id": 123
}
```
