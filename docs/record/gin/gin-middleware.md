## gin-middleware

知识要点：

中间件的编写

代码：

```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
        fmt.Println("before middleware")
        //设置request变量到Context的Key中,通过Get等函数可以取得
        c.Set("name", "username-xxx")
        c.Set("user", "hanru")
        //发送request之前
        c.Next()

        //发送requst之后

        // 这个 c.Write 是 ResponseWriter,我们可以获得状态等信息
        status := c.Writer.Status()
        fmt.Println("after middleware,", status)
        t2 := time.Since(t)
        fmt.Println("time:", t2)
    }
}

func main() {
    router := gin.Default()
    router.Use(MiddleWare())
    {
        router.GET("/middleware", func(c *gin.Context) {
            // 获取gin上下文中的变量
            // 另外，如果没有注册就使用MustGet方法读取c的值将会抛错，可以使用Get方法取而代之。
            // 上面的注册装饰方式，会让所有下面所写的代码都默认使用了router 的注册过的中间件。
            name := c.MustGet("name").(string)
            req, _ := c.Get("name")
            fmt.Println("name:", name)
            c.JSON(http.StatusOK, gin.H{
                "middile_name": name,
                "request":      req,
            })
        })

        router.GET("/before", MiddleWare(), func(c *gin.Context) {
            username := c.MustGet("name").(string)
            c.JSON(http.StatusOK, gin.H{
                "middile_request": username,
            })
        })
    }

    // 模拟私有数据
    var secrets = gin.H{
        "hanru":     gin.H{"email": "hanru@163.com", "phone": "123433"},
        "wangergou": gin.H{"email": "wangergou@example.com", "phone": "666"},
        "ruby":      gin.H{"email": "ruby@guapa.com", "phone": "523443"},
    }
    authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
        // 已授权的账号密码
        "hanru":     "hanru123",
        "wangergou": "1234",
        "ruby":      "hello2",
        "lucy":      "4321",
    }))

    // 在浏览器打开这个地址：http://127.0.0.1:8080/admin/secrets，按照提示输入上面的账号密码
    authorized.GET("/secrets", func(c *gin.Context) {
        fmt.Println("======>>>>>>>>>>>>")
        // 获取提交的用户名（AuthUserKey）
        user := c.MustGet(gin.AuthUserKey).(string)
        if secret, ok := secrets[user]; ok {
            c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
        } else {
            c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
        }
    })

    // Listen and serve on 0.0.0.0:8080
    router.Run(":8080")
}
```

测试：

```http
@base = 127.0.0.1:8080

### 
GET http://{{base}}/middleware HTTP/1.1

###
GET http://{{base}}/before HTTP/1.1
###
```
