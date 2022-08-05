## gin-template-render 模板渲染



知识要点：





操作步骤：



- 本地新建 `templates` 目录，添加以下文件
  
  - `templates/index.tmpl` 代码为：
  
  ```tmpl
  <html>
      <h1>
         Page-1 {{ .title }}
      </h1>
  </html>
  ```
  
  
  
  - `templates/posts/index.tmpl`
  
  ```tmpl
  {{ define "posts/index.tmpl" }}
  <html>
      <h1>
         posts {{ .title }}
      </h1>
  </html>
  {{ end }}
  ```
  
  
  
  - `templates/user/index.tmpl`
  
  ```tmpl
  {{ define "user/index.tmpl" }}
  <html>
      <h1>
         user {{ .title }}
      </h1>
  </html>
  {{ end }}
  ```
  
  

代码：



```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//--------- 单个 index 模板 -----------------
	// router.LoadHTMLGlob("templates")
	// router.GET("/index", func(c *gin.Context) {
	// 	//根据完整文件名渲染模板，并传递参数
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Hello, World!",
	// 	})
	// })

	// -------- 多个 index 模板 -----------------
	// 多个 index 模板，使用 * 匹配多个层级，index 同名也没关系
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/posts/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Hello, Posts!",
		})
	})

	router.GET("/user/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
			"title": "Hello, User!",
		})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

```



测试：



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




