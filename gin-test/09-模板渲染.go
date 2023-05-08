package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	testHttp := gin.Default()
	/*
		//1.加载模板 参数可以是多个
		testHttp.LoadHTMLFiles("./templates/index.tmpl", "./templates/user.tmpl")
		testHttp.GET("/index", func(ctx *gin.Context) {
			//加载已经创建好的tmpl模板，并填充值
			ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "hello",
			})
		})
		//2.多模板渲染
		testHttp.GET("/user", func(ctx *gin.Context) {
			//加载已经创建好的tmpl模板，并填充值
			ctx.HTML(http.StatusOK, "user.tmpl", gin.H{
				"title": "hello2",
			})
		})
	*/

	//3.另一种多模板渲染方式 (仅同目录)
	//testHttp.LoadHTMLGlob("./templates/**")
	//4.另一种多模板渲染方式 (仅子目录)
	testHttp.LoadHTMLGlob("./templates/**/*")

	testHttp.GET("/index", func(ctx *gin.Context) {
		//加载已经创建好的tmpl模板，并填充值
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "hello",
		})
	})
	testHttp.GET("/user", func(ctx *gin.Context) {
		//加载已经创建好的tmpl模板，并填充值
		ctx.HTML(http.StatusOK, "user.tmpl", gin.H{
			"title": "hello2",
		})
	})

	err := testHttp.Run(":8088")
	if err != nil {
		fmt.Println("运行失败")
	}

}
