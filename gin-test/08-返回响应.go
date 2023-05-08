package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type XmlUser struct {
	Id   int64  `xml:"id"`
	Name string `xml:"name"`
}

func main() {
	testHttp := gin.Default()
	testHttp.GET("/user/save", func(ctx *gin.Context) {
		//string的返回响应
		//ctx.String(http.StatusOK, "this is %s", "test page")

		//JSON返回响应
		//ctx.JSON(http.StatusOK, gin.H{
		//	"success": true,
		//})

		/*
			//	使用xml形式
			//ctx.XML(http.StatusOK, gin.H{
			//	"success": true,
			//})
		*/

		/*
			//	使用结构体返回xml
			ctx.XML(http.StatusOK, XmlUser{
				Id:   123,
				Name: "测试",
			})
		*/

		// 返回文件
		//ctx.File("./测试.png")
		//返回文件起别名
		//ctx.FileAttachment("./测试.png", "1.png")

		//	返回请求头
		//ctx.Header("test", "ms")

		//	重定向
		//ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")

		//	返回ymal
		ctx.YAML(http.StatusOK, gin.H{
			"name": "ms",
			"age":  19,
		})
	})

	err := testHttp.Run(":8087")
	if err != nil {
		fmt.Println("运行失败")
	}

}
