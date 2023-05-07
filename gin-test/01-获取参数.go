package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	testHttps := gin.Default()
	//http://localhost:8080/user/save?id=11&name=zhaoyang
	testHttps.GET("/user/save", func(ctx *gin.Context) {
		//获取uri中的指定参数
		id := ctx.Query("id")
		name := ctx.Query("name")
		//获取uri可能的参数,ok是判断是否存在这个参数。因为有些参数是存在的，但是值为nil
		//address, ok := ctx.GetQuery("address")
		//这里是如果不存在这个参数，那么就赋予一个默认值
		address := ctx.DefaultQuery("address", "shanghai")
		ctx.JSON(200, gin.H{
			"id":      id,
			"name":    name,
			"address": address,
		})
	})

	err := testHttps.Run(":8080")
	if err != nil {
		log.Fatalln("运行失败！")
	}
}
