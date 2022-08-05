## gin-response-type

**知识要点：**

- `c.JSON`

- `c.XML`

- `c.YAML`

- `c.ProtoBuf`

Code:

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
    router := gin.Default()

    // gin.H is a shortcut for map[string]interface{}
    router.GET("/json-gin", func(c *gin.Context) {
        // c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
        c.JSON(http.StatusOK, gin.H{"code": 0, "data": nil})
    })

    router.GET("/json", func(c *gin.Context) {
        // You also can use a struct
        var msg struct {
            Name    string `json:"userName"`
            Message string `json:"msg"`
            Number  int    `json:"num"`
        }
        msg.Name = "Tony"
        msg.Message = "理发"
        msg.Number = 100
        // 注意 msg.Name 变成了 "user" 字段
        // 以下方式都会输出 :   {"userName": "Tony", "msg": "理发", "num": 100}
        c.JSON(http.StatusOK, msg)
    })

    router.GET("/json-xml", func(c *gin.Context) {
        c.XML(http.StatusOK, gin.H{"user": "Tony", "message": "hey", "status": http.StatusOK})
    })

    router.GET("/json-yml", func(c *gin.Context) {
        c.YAML(http.StatusOK, gin.H{"message": "Tony", "status": http.StatusOK})
    })

    router.GET("/protobuf", func(c *gin.Context) {
        reps := []int64{int64(1), int64(2)}
        label := "test"
        // The specific definition of protobuf is written in the testdata/protoexample file.
        data := &protoexample.Test{
            Label: &label,
            Reps:  reps,
        }
        // Note that data becomes binary data in the response
        // Will output protoexample.Test protobuf serialized data
        c.ProtoBuf(http.StatusOK, data)
    })

    // Listen and serve on 0.0.0.0:8080
    router.Run(":8080")
}
```

测试

```http
@base = 127.0.0.1:8080

### c.json 响应 gin.H 方式返回
GET http://{{base}}/json-gin HTTP/1.1

### c.json 响应  自定义对象的方式返回
GET http://{{base}}/json HTTP/1.1

### c.xml 响应 的方式返回
GET http://{{base}}/json-xml HTTP/1.1

### c.yml 响应 的方式返回
GET http://{{base}}/json-yml HTTP/1.1

### c.protobuf 响应 的方式返回
GET http://{{base}}/protobuf HTTP/1.1
```
