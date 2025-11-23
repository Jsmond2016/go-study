# Gin 学习记录

> 注意：此目录下的文档已整合到 [Web 开发](../web-development/) 模块中

## 📚 已整合的文档

以下文档内容已整合到对应的正式文档中：

- **gin-begin** → [Gin 基础](../web-development/02-gin-basics.md)
- **gin-router** → [Gin 路由](../web-development/03-gin-routing.md)
- **gin-router-params** → [Gin 路由](../web-development/03-gin-routing.md)
- **gin-router-group** → [Gin 路由](../web-development/03-gin-routing.md)
- **gin-router-params-bind** → [Gin 路由](../web-development/03-gin-routing.md)
- **gin-response-type** → [Gin 基础](../web-development/02-gin-basics.md)
- **gin-template-render** → [Gin 模板](../web-development/05-gin-template.md)
- **gin-static** → [Gin 模板](../web-development/05-gin-template.md)
- **gin-sync-async** → [Gin 基础](../web-development/02-gin-basics.md)
- **gin-middleware** → [Gin 中间件](../web-development/04-gin-middleware.md)
- **gin-mysql** → [Gin 数据库操作](../web-development/10-gin-database.md)

## 🎯 推荐学习路径

请按照以下顺序学习 Gin 框架：

1. [Gin 基础](../web-development/02-gin-basics.md) - 安装、基本使用、响应类型、异步处理
2. [Gin 路由](../web-development/03-gin-routing.md) - 路由配置、参数获取、路由组、参数绑定
3. [Gin 中间件](../web-development/04-gin-middleware.md) - 中间件开发和使用
4. [Gin 模板](../web-development/05-gin-template.md) - 模板渲染、静态文件服务
5. [Gin 数据库操作](../web-development/10-gin-database.md) - MySQL CRUD 操作
6. [数据验证](../web-development/06-gin-validation.md) - 请求验证和绑定
7. [认证授权](../web-development/07-gin-auth.md) - JWT、Session
8. [REST API 设计](../web-development/08-rest-api.md) - API 设计最佳实践

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
