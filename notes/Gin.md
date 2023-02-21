### Gin

#### [安装](https://pkg.go.dev/github.com/gin-gonic/gin#readme-installation)

```bash
go get -u github.com/gin-gonic/gin
```

#### 快速启动

```GO
package main

import "github.com/gin-gonic/gin"

func main() {
	// Gin
	ginServer := gin.Default()

	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "Hello, Gin."})
	})

	ginServer.Run(":8888") // 注意这里的冒号

}
```

#### 请求与响应
- [x] ApiPost / Postman

```GO
ginServer.GET("/hello", func(context *gin.Context) {
    context.JSON(200, gin.H{"msg": "Hello, Gin."})
})
```

```GO
```