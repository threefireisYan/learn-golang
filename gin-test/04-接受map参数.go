package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User3 struct {
	Id         int64             `form:"id"`
	Name       string            `form:"name"`
	Address    string            `form:"address"`
	AddressMap map[string]string `form:"addressMap" binding:"required"`
}

func main() {
	testhttp := gin.Default()

	//http://localhost:8083/user/save
	testhttp.GET("/user/save", func(ctx *gin.Context) {
		addressMap := ctx.QueryMap("addressMap")
		//不支持直接使用绑定结构体来获取map参数
		var user User3
		ctx.ShouldBindQuery(&user) //这种形式是不能够获取到的
		//但是可以先使用QueryMap获取到参数后赋值给结构体
		user.AddressMap = addressMap
		ctx.JSON(200, user)

	})

	err := testhttp.Run(":8083")
	if err != nil {
		fmt.Println("运行失败")
	}
}
