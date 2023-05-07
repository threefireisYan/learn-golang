package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Id      int64  `form:"id"`
	Name    string `form:"name"`
	Address string `form:"address" binding:"required"` //表示字段为必须字段
}

func main() {
	testHttps := gin.Default()
	//http://localhost:8081/user/save?id=11&name=zhaoyang
	testHttps.GET("/user/save", func(ctx *gin.Context) {
		//使用结构体来获取uri中的参数
		var user User
		//一般不使用BindQuery,如果uri中没有参数，会返回400，并且参数不是JSON格式
		//err := ctx.BindQuery(&user)
		//一般使用ShouldBindQuery，不会报错，参数是JSON格式
		err := ctx.ShouldBindQuery(&user)
		if err != nil {
			fmt.Println(err)
		}
		ctx.JSON(200, user)
	})
	err := testHttps.Run(":8081")
	if err != nil {
		log.Fatalln("运行失败！")
	}
}
