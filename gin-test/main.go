package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	testhttps := gin.Default()
	//crul http://localhost:8080/hello  GET 获取到 json 返回值 {"name":"hello world"}
	testhttps.GET("/hello", func(ctx *gin.Context) {
		//可返回数据 数组 map list struct机构提
		ctx.JSON(200, gin.H{
			"name": "hello world",
		})
	})
	//启动服务  8080前面的冒号不能省略
	err := testhttps.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
