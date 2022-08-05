## gin-sync-async 异步处理

> `gin-demo\chapter-04\exercise-4\main.go`

知识要点：

处理一些耗时任务时，使用异步操作



代码：



```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//1. 异步，类似 setTimeout，先返回结果， go func 语句后台继续执行
	router.GET("/upload-async", func(c *gin.Context) {
		// goroutine 中只能使用只读的上下文 c.Copy()
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			// 注意使用只读上下文
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
		c.JSON(http.StatusOK, "异步请求成功，请稍后尝试")
	})
	//2. 同步，要等待 处理时间，处理完成后返回
	router.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		// 注意可以使用原始上下文
		log.Println("Done! in path " + c.Request.URL.Path)
		c.JSON(http.StatusOK, "同步处理中，耗时 5s ")
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

```

测试：



```http
@base = 127.0.0.1:8080

### 
GET http://{{base}}/upload-async HTTP/1.1

###
GET http://{{base}}/sync HTTP/1.1
```


