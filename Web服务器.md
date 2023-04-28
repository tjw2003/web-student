

# Web服务器

### Gin包

`Gin`是一个用`Golang`编写的`Web`框架， 是一个基于`Radix`树的路由，小内存占用，没有反射，可预测的`API`框架，速度挺高了近40倍，支持中间件、路由组处理、JSON等多方式验证、内置了`JSON`、`XML`、`HTML`等渲染。是一个易于使用的`API`框架

```
go mod init <名>//建立一个模块，方便应用各种包

//得到gin包的命令
go get -u github.com/gin-gonic/gin
```

### 一个简单的web服务器实例

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```



### 路由route

简单来说，路由就是URL到函数的映射。route就是一条路由，它将一个URL路径和一个函数进行映射。

例如：

```
/users        ->  getAllUsers()
/users/count  ->  getUsersCount()
```

这就是两条路由，当访问 /users 的时候，会执行 getAllUsers() 函数；当访问 /users/count 的时候，会执行 getUsersCount() 函数。

而 router 可以理解为一个容器，或者说一种机制，它管理了一组 route。简单来说，route 只是进行了URL和函数的映射，而在当接收到一个URL之后，去路由映射表中查找相应的函数，这个过程是由 router 来处理的。

对于服务器来说，当接收到客户端发来的HTTP请求，会根据请求的URL，来找到相应的映射函数，然后执行该函数，并将函数的返回值发送给客户端。

```
app.get('/', (req, res) => {
  res.sendFile('index')
})

app.get('/users', (req, res) => {
  db.queryAllUsers()
    .then(data => res.send(data))
})
```

- 这是两条路由
- 当访问 / 的时候，会返回 index 页面
- 当访问 /users 的时候，会从数据库中取出所有用户数据并返回

在gin包中

```goapi := r.Group("api")
api := r.Group("api")
	{
		api.GET("user", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"name": "tjw"})
		})

		api.GET("contract", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"date": "tody"})
		})
	}

```

api/user 返回 name: tjw

api/contract 返回 date : tody

### 文件服务器

主要用到的方法是http包的FileServer

```go
http.FileServer(http.Dir("文件名"))

http.Handle("/", http.FileServer(http.Dir("./")))
go http.ListenAndServe(":8080", nil)//携程


http.ListenAndSrver(":8080", http.FileServer(http.Dir("./")))
```

```
main.css//如果不用文件服务器，r.GET("/", func(ctx *gin.Context) {ctx.File("")})，仅作这个操作main.css不会找到
web.html
```

文件服务器会与之前的api产生冲突，所以要用到中间件

### 中间件（middleware）

```go
func middleware(ctx *gin.Context) { //中间件
	path := ctx.Request.URL.Path//得到路径

	if strings.HasPrefix(path, "/api") {
		ctx.Next()//继续向下运行
	} else {
		http.FileServer(http.Dir("./frontend/")).ServeHTTP(ctx.Writer, ctx.Request)
		ctx.Abort()//停止
	}
}

r.Use(middleware)
```

### 数据库使用

Go 没有内置的驱动支持任何的数据库，但是 Go 定义了 database/sql 接口，用户可以基于驱动接口开发相应数据库的驱动

Go 实现的支持 PostgreSQL 的驱动

```go
github.com/jackc/pgx/v5/stdlib
```

数据库连接

```go
func openDatabase() {
	db, err := sql.Open("pgx", os.Getenv("postgres://postgres:@localhost:5432/postgres"))
	if err != nil {
		log.Panicf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()
}
```

### 测试文件

```go
package database

import (
	"testing"
)

func TestOpenDatabase(t *testing.T) {
	openDatabase()
	t.Log("sdffsdf")
}
```

在目录下命令go test,正确则log不会有输出

go -test.v test,log会有输出

