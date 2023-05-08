package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	testhttp := gin.New()
	//通过user 设置全局中间件
	//设置日志中间件，主要用于打印请求日志
	testhttp.Use(gin.Logger())
	//	设置Recovery中间件，主要用于拦截paic错误，不至于导致进程崩溃
	testhttp.Use(gin.Recovery())
	testhttp.GET("/test", func(ctx *gin.Context) {
		panic(errors.New("test error"))
	})
	err := testhttp.Run(":8095")
	if err != nil {
		fmt.Println("运行失败")
	}

}
