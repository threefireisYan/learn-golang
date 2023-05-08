package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	testHttp := gin.Default()

	//自定义一定要在加载模板前面
	testHttp.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//加载静态文件
	//testHttp.Static("/css", "./static/css")

	//加载子目录所有模板
	testHttp.LoadHTMLGlob("./templates/**/*")

	testHttp.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "<a href = 'http://www.baidu.com'>hello templates</a>",
		})
	})

	err := testHttp.Run(":8089")
	if err != nil {
		fmt.Println("运行失败")
	}
}
