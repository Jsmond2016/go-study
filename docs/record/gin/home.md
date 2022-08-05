## 导航页

> 代码参考：[gin_demo](https://github.com/Jsmond2016/gin_demo)

学习顺序

- gin-begin

- gin-ping

- gin-router

- gin-router-params

- gin-router-group

- gin-router-params-bind

- gin-response-type

- gin-template-render

- gin-static

- gin-sync-async

- gin-middleware

- gin-mysql

## 其他准备：

学习资料：https://www.chaindesk.cn/witbook/19/330

- 创建数据库

步骤：

```bash
docker pull mysql:latest



docker images



docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
```

使用mysql [可视化工具 heidisql](https://www.heidisql.com/download.php)，查看

mysql 连接的账号密码：

```
root/123456
```

- 建表

```sql
CREATE TABLE IF NOT EXISTS `user_info`(

   `id` INT UNSIGNED AUTO_INCREMENT,

   `username` VARCHAR(100) NOT NULL,

   `password` VARCHAR(40) NOT NULL,

   PRIMARY KEY ( `id` )

)ENGINE=InnoDB DEFAULT CHARSET=UTF8;



USE test;
```
