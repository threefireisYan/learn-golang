package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	testHttp := gin.Default()

	testHttp.POST("/user/save", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			log.Fatalln(err)
		}
		values := form.Value
		files := form.File
		for _, fileArray := range files {
			for _, v := range fileArray {
				ctx.SaveUploadedFile(v, "./"+v.Filename)
			}
		}
		ctx.JSON(200, values)
	})

	err := testHttp.Run(":8086")
	if err != nil {
		fmt.Println("运行失败")
	}
}
