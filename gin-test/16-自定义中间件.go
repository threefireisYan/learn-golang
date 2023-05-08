package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("before action")
		t := time.Now()
		//可以通过上下文对象，设置一些依附在上下文对象里面的键值数据
		ctx.Set("example", "123")

		ctx.Next()
		value, _ := ctx.Get("aaa")
		//在这里可以处理请求返回给用户之前的逻辑
		latency := time.Since(t)
		log.Println(latency, value)
		fmt.Println("after action")
		//查询请求状态
		status := ctx.Writer.Status()
		log.Println(status)

	}
}

func main() {
	testhttp := gin.New()
	//通过user 设置全局中间件
	//设置日志中间件，主要用于打印请求日志
	testhttp.Use(gin.Logger())
	//	设置Recovery中间件，主要用于拦截paic错误，不至于导致进程崩溃
	testhttp.Use(gin.Recovery())
	testhttp.Use(Logger())
	testhttp.GET("/test", func(ctx *gin.Context) {
		//panic(errors.New("test error"))
		value, exists := ctx.Get("example")
		fmt.Println("example value", value, "exists:", exists)
		ctx.Set("aaa", "bbb")
		ctx.JSON(200, "success")
	})
	err := testhttp.Run(":8095")
	if err != nil {
		fmt.Println("运行失败")
	}
}
