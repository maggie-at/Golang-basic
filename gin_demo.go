package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func MyMidHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 通过自定义中间件设置值, 后续处理只要调用这个中间件, 都可以拿到这里的值
		context.Set("userSession", "userSession~~~")
		context.Next() // 放行
	}
}
func GinMain() {
	// Gin
	ginServer := gin.Default()

	// 设置favicon图标
	ginServer.Use(favicon.New("./templates/favicon.ico"))

	// 注册中间件
	ginServer.Use(MyMidHandler())

	// RestFul风格接口					// 中间件拦截器
	ginServer.GET("/helloGet", MyMidHandler(), func(context *gin.Context) { // 函数式编程
		userSession := context.MustGet("userSession").(string)
		log.Println(userSession)
		context.JSON(200, gin.H{"msg": "[GET] Hello, Gin."}) // gin.H(): 一个map
	})

	ginServer.POST("/helloPost", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "[POST] Hello, Gin."})
	})

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*")
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{ // 同时还可以传数据
			"msg": "A static HTML page",
		})
	})

	// 接收前端传递的参数
	// 方式一: /user/info?userid=1&username=alan   +	context.Query("key")
	ginServer.GET("/user/info", func(context *gin.Context) {
		user_id := context.Query("userid")
		user_name := context.Query("username")
		// 还可以返回参数: JSON serializes the given struct as JSON into the response body.
		context.JSON(http.StatusOK, gin.H{
			"userid":   user_id,
			"username": user_name,
		})
	})
	// 方式二: /user/info/1/alan		+	context.Param("key")
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userId := context.Param("userid")
		userName := context.Param("username")
		// 还可以返回参数
		context.JSON(http.StatusOK, gin.H{
			"userid":   userId,
			"username": userName,
		})
	})

	// 获取前端传递的JSON, 并进行序列化
	ginServer.POST("/json", func(context *gin.Context) {
		data, _ := context.GetRawData()

		var mp map[string]interface{} // GO语言的Object可以用空接口表示
		_ = json.Unmarshal(data, &mp)
		//for k, v := range mp {
		//	fmt.Println(k, v)
		//}
		context.JSON(http.StatusOK, mp)
	})

	// 处理表单
	ginServer.POST("/user/add", func(context *gin.Context) {
		// 获取表单参数
		userName := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{
			"status":   "active",
			"userName": userName,
			"password": password,
		})
	})

	// 路由
	// 1. 重定向 301
	ginServer.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	// 2. 404
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "notfound.html", nil)
	})

	// 3. 路由组
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.POST("/login")
		userGroup.POST("/logout")
	}
	orderGroup := ginServer.Group("/order")
	{
		orderGroup.GET("/add")
		orderGroup.POST("/modify")
	}

	ginServer.Run(":8888") // 注意这里的冒号
}
