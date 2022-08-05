## gin-router-params

**知识要点：** 

> `c *gin.Context`

- `c.Query`

- `c.DefaultQuery`

- `c.PostForm`

- `c.DefaultPostForm`

- `c.FormFile`

- `c.PostFormArray`

- `c.MultipartForm`

Code:

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/get", func(c *gin.Context) {
        name := c.DefaultQuery("name", "匿名者") //可设置默认值
        //nickname := c.Query("nickname") // 是 c.Request.URL.Query().Get("nickname") 的简写
        c.String(http.StatusOK, fmt.Sprintf("Hello %s ", name))
    })

    //form
    router.POST("/form", func(c *gin.Context) {
        // type1 := c.DefaultPostForm("type", "alert") //可设置默认值
        username := c.PostForm("username")
        password := c.PostForm("password")

        //hobbys := c.PostFormMap("hobby")
        //hobbys := c.QueryArray("hobby")
        hobbys := c.PostFormArray("hobby")

        c.String(http.StatusOK, fmt.Sprintf("username is %s, password is %s,hobby is %v", username, password, hobbys))

    })

    router.POST("/upload", func(c *gin.Context) {
        // single file
        file, _ := c.FormFile("file")
        log.Println(file.Filename)

        // Upload the file to specific dst.
        c.SaveUploadedFile(file, file.Filename)

        /*
            也可以直接使用io操作，拷贝文件数据。
            out, err := os.Create(filename)
            defer out.Close()
            _, err = io.Copy(out, file)
        */

        c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
    })

    router.MaxMultipartMemory = 8 << 20 // 8 MiB
    //router.Static("/", "./public")
    router.POST("/uploads", func(c *gin.Context) {

        // Multipart form
        form, err := c.MultipartForm()
        if err != nil {
            c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
            return
        }
        files := form.File["files"]

        for _, file := range files {
            if err := c.SaveUploadedFile(file, file.Filename); err != nil {
                c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
                return
            }
        }

        c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files ", len(files)))
    })

    router.Run()
}
```

测试代码：

- 前端

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>登录</title>
</head>
<body>
    <h2>表单数据</h2>
    <form action="http://127.0.0.1:8080/form" method="post" enctype="application/x-www-form-urlencoded">
        用户名：<input type="text" name="username">
        <br>
        密&nbsp&nbsp&nbsp码：<input type="password" name="password">
        <br>
        兴&nbsp&nbsp&nbsp趣：
        <input type="checkbox" value="girl" name="hobby">女人
        <input type="checkbox" value="game" name="hobby">游戏
        <input type="checkbox" value="money" name="hobby">金钱
        <br>
        <input type="submit" value="登录">
    </form>
    <hr>
    <h2>单文件上传</h2>
    <form action="http://127.0.0.1:8080/upload" method="post" enctype="multipart/form-data">
        头像：
        <input type="file" name="file">
        <br>
        <input type="submit" value="提交">
    </form>
    <h2>多文件上传</h2>
    <form action="http://127.0.0.1:8080/uploads" method="post" enctype="multipart/form-data">
        Files: <input type="file" name="files" multiple><br><br>
        <input type="submit" value="提交">
    </form>
</body>
</html>
```

接口测试：

```http
@base = 127.0.0.1:8080

### 

GET http://{{base}}/get?name=测试人员&num=2 HTTP/1.1

###
POST http://{{base}}/form?name=测试人员&num=2 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

username=Tony
&password=123455
&hobby="basketball"
&hobby="running"
```
