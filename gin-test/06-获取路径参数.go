package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User6 struct {
	Id         int64             `json:"id,omitempty" uri:"id"`
	Name       string            `json:"name,omitempty" uri:"name"`
	Address    []string          `json:"address,omitempty"`
	AddressMap map[string]string `json:"address_map,omitempty"`
}

func main() {
	testHttps := gin.Default()
	testHttps.POST("/user/save/:id/:name", func(ctx *gin.Context) {
		//直接使用获取参数形式
		//所以一般使用直接获取参数的形式来对uri中的参数进行获取使用
		id := ctx.Param("id")
		name := ctx.Param("name")
		ctx.JSON(200, gin.H{
			"id":   id,
			"name": name,
		})

		//	使用结构体获取路径,这种方式如果在uri中前一个参数获取有问题，那么后面的参数都将获取不到
		//	比如/user/save/:id/:name  id在结构体中是一个int64的类型，在uri中/user/save/123cs/test  那么是无法获取到id=123cs以及name=test的
		//var user User6
		//ctx.ShouldBindUri(&user)
		//ctx.JSON(200, user)
	})

	err := testHttps.Run(":8085")
	if err != nil {
		fmt.Println("运行失败")
	}
}
