package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 结构体内的变量名称一定要首字母大写，否则权限是私有权限
// 仅在结构体内或结构体内的函数方法能够直接访问，在外部调用结构体变量无法访问赋值
type user2 struct {
	Id      int64    `form:"id"`
	Name    string   `form:"name"`
	Address []string `form:"address" binding:"required"`
}

func main() {
	testHttps := gin.Default()
	//http://localhost:8082/user/save?address="beijing"&address="shanghai"
	testHttps.GET("/user/save", func(ctx *gin.Context) {
		//指定获取address参数
		//address := ctx.QueryArray("address")
		//使用确认是否含有address参数
		//address, _ := ctx.GetQueryArray("address")
		//数组不存在default的函数
		//同样可以使用结构体进行接受参数
		var user user2
		ctx.ShouldBindQuery(&user)
		ctx.JSON(200, user)
		//ctx.JSON(200, address)
	})
	err := testHttps.Run(":8082")
	if err != nil {
		fmt.Println("运行失败！")
	}
}
