package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	testhttps := gin.Default()
	//crul http://localhost:8080/hello  GET 获取到 json 返回值 {"name":"hello world"}
	//testhttps.GET("/hello", func(ctx *gin.Context) {
	//	//可返回数据 数组 map list struct机构提
	//	ctx.JSON(200, gin.H{git
	//		"name": "hello world",
	//	})
	//})
	/*
		//查询
		testhttps.GET("/user", func(ctx *gin.Context) {
			ctx.JSON(200, "get")
		})
		//增加
		testhttps.POST("/user", func(ctx *gin.Context) {
			ctx.JSON(200, "add")
		})
		//删除
		testhttps.DELETE("/user", func(ctx *gin.Context) {
			ctx.JSON(200, "delete")
		})
		//修改
		testhttps.PUT("/user", func(ctx *gin.Context) {
			ctx.JSON(200, "update")
		})
	*/
	/*
		//任意类型的http uri请求都支持 （这里的relativePath都是静态路径参数）
		testhttps.Any("/any", func(ctx *gin.Context) {
			ctx.JSON(200, "update")
		})
	*/
	// ctx.Param("")获取的是热辣tivePath中的匹配字段
	/*
		//任意类型的http uri请求都支持 （这里的relativePath是动态参数）
		testhttps.Any("/user/find/:id", func(ctx *gin.Context) {
			ctx.JSON(200, ctx.Param("id"))
		})
	*/
	/*
		//任意类型的http uri请求都支持 （这里的relativePath是模糊匹配参数 或叫 做通配符）
		testhttps.Any("/user/*path", func(ctx *gin.Context) {
			ctx.JSON(200, ctx.Param("path"))
		})
	*/

	//分组路由
	//将下列请求返回体进行分组
	//testhttps.GET("v1/user/save", func(ctx *gin.Context) {
	//	ctx.JSON(200, ctx.Param("path"))
	//})
	//
	//testhttps.GET("v2/user/save", func(ctx *gin.Context) {
	//	ctx.JSON(200, ctx.Param("path"))
	//})
	//
	//testhttps.GET("/user/find", func(ctx *gin.Context) {
	//	ctx.JSON(200, ctx.Param("path"))
	//})
	//版本分组
	v1 := testhttps.Group("/v1")
	{
		v1.GET("find", func(ctx *gin.Context) {
			ctx.JSON(200, "v1 find")
		})
		v1.GET("save", func(ctx *gin.Context) {
			ctx.JSON(200, "v1 save")
		})
	}
	v2 := testhttps.Group("/v2")
	{
		v2.GET("find", func(ctx *gin.Context) {
			ctx.JSON(200, "v2 find")
		})
		v2.GET("save", func(ctx *gin.Context) {
			ctx.JSON(200, "v2 save")
		})
	}

	//启动服务  8080前面的冒号不能省略
	err := testhttps.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
