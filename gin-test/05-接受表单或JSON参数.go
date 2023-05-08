package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
type User5 struct {
	Id         int64             `form:"id"`
	Name       string            `form:"name"`
	Address    []string          `form:"address"`
	AddressMap map[string]string `form:"addressMap" binding:"required"`
}
*/

// 如果使用JSON接受值，那么不应用form，应用json
type User5 struct {
	Id         int64             `json:"id"`
	Name       string            `json:"name"`
	Address    []string          `json:"addressss"`
	AddressMap map[string]string `json:"addressMap" binding:"required"`
}

func main() {
	testHttp := gin.Default()

	testHttp.POST("/user/save", func(ctx *gin.Context) {
		/*
			id := ctx.PostForm("id")
			name := ctx.PostForm("name")
			address := ctx.PostFormArray("address")
			addressMap := ctx.PostFormMap("addressMap")
			ctx.JSON(200, gin.H{
				"id":         id,
				"name":       name,
				"address":    address,
				"addressMap": addressMap,
			})
			以上与get方式获取对应参数方式相同
		*/
		/*
			//这里的绑定结构体，与get不一样，是直接ShouldBind,而不是ShouldBindQuery
			var user User5
			ctx.ShouldBind(&user)
			addressMap := ctx.PostFormMap("addressMap")
			user.AddressMap = addressMap
			ctx.JSON(200, user)
		*/
		//若前端传来的是JSON格式的参数，那么可以直接使用ShouldBindJSON绑定结构体
		var user User5
		ctx.ShouldBindJSON(&user)
		ctx.JSON(200, user)

	})

	err := testHttp.Run(":8084")
	if err != nil {
		fmt.Println("运行失败")
	}
}
