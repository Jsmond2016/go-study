gin-mysql



知识要点：对 MySQL 的 crud 操作

- MySql 的初始化和释放连接
- 查询
  - 单个查询
  - 多个查询
  - 模糊查询
- 新增
  - 单个新增
  - 多个新增
- 更新
  - 单个更新
  - 多个更新
- 删除
  - 单个删除
  - 多个删除





Code：

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//定义User类型结构
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func initDB() *sql.DB {
	db, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	return db
}

//定义一个getALL函数用于回去全部的信息
func getAll() (users []User, err error) {

	//1.操作数据库
	db := initDB()
	//错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	//推迟数据库连接的关闭
	defer db.Close()

	//2.查询
	rows, err := db.Query("SELECT id, username, password FROM user_info")
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var user User
		//遍历表中所有行的信息
		rows.Scan(&user.Id, &user.Username, &user.Password)
		//将user添加到users中
		users = append(users, user)
	}
	//最后关闭连接
	defer rows.Close()
	return
}

//插入数据
func add(user User) (Id int, err error) {

	//1.操作数据库
	db := initDB()
	//错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	//推迟数据库连接的关闭
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO user_info(username, password) VALUES (?, ?)")
	if err != nil {
		return
	}
	//执行插入操作
	rs, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return
	}
	//返回插入的id
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	//将id类型转换
	Id = int(id)
	defer stmt.Close()
	return
}

//修改数据
func update(user User) (rowsAffected int64, err error) {

	//1.操作数据库
	db := initDB()
	//错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	//推迟数据库连接的关闭，确保 db 不为 Nil
	defer db.Close()
	stmt, err := db.Prepare("UPDATE  user_info SET username=?, password=? WHERE id=?")
	if err != nil {
		return
	}
	//执行修改操作
	rs, err := stmt.Exec(user.Username, user.Password, user.Id)
	if err != nil {
		return
	}
	//返回插入的id
	rowsAffected, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	return
}

//通过id删除
func del(id int) (rows int, err error) {
	// 1.操作数据库
	db := initDB()
	//错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	//推迟数据库连接的关闭
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM user_info WHERE id=?")
	if err != nil {
		log.Fatalln(err)
	}

	rs, err := stmt.Exec(id)
	if err != nil {
		log.Fatalln(err)
	}
	//删除的行数
	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rows = int(row)
	return
}

func main() {
	//创建一个路由Handler
	router := gin.Default()

	//get方法的查询
	router.GET("/user", func(c *gin.Context) {
		users, err := getAll()
		if err != nil {
			log.Fatal(err)
		}
		//H is a shortcut for map[string]interface{}
		c.JSON(http.StatusOK, gin.H{
			"result": users,
			"count":  len(users),
		})
	})

	router.POST("/add", func(c *gin.Context) {
		var u User
		err := c.Bind(&u)
		if err != nil {
			log.Fatal(err)
		}
		Id, err := add(u)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("id=", Id)
		fmt.Println("===========================================")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s 插入成功, id 为 %d", u.Username, Id),
		})
	})

	// 利用put方法修改数据
	router.PUT("/update", func(c *gin.Context) {
		var u User
		err := c.Bind(&u)
		if err != nil {
			log.Fatal(err)
		}
		num, _ := update(u)
		fmt.Print("num=", num)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("修改id: %d 成功", u.Id),
		})
	})

	//利用DELETE请求方法通过id删除
	router.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")

		Id, err := strconv.Atoi(id)

		if err != nil {
			log.Fatalln(err)
		}
		rows, err := del(Id)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("delete rows ", rows)

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted user: %s", id),
		})
	})

	router.Run(":8080")
}

```





测试：



```http
@base = 127.0.0.1:8080

### 
GET http://{{base}}/user HTTP/1.1


###
POST http://{{base}}/add HTTP/1.1
Content-Type: application/json

{
  "username": "Delete test",
  "password": "test"
}

###
PUT http://{{base}}/update HTTP/1.1
Content-Type: application/json

{
  "id": 1,
  "username": "Tony",
  "password": "3333333"
}
###
DELETE http://{{base}}/delete/3 HTTP/1.1


```
