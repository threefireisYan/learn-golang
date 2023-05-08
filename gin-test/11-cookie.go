package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	testHttp := gin.Default()

	testHttp.GET("/cookie", func(ctx *gin.Context) {
		//设置cookie
		ctx.SetCookie("site_cookie", "cookievalue", 3600, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	testHttp.GET("/readCookie", func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("site_cookie")
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, gin.H{
				"error": "Cookie not found",
			})
			return
		}

		value := cookie.Value
		ctx.JSON(http.StatusOK, gin.H{
			"cookie": value,
		})
	})

	//删除cookie
	testHttp.GET("/deleteCookie", func(ctx *gin.Context) {
		ctx.SetCookie("site_cookie", "cookievalue", -1, "/", "localhost", false, true)
	})

	err := testHttp.Run(":8090")
	if err != nil {
		log.Fatalln(err)
	}
}
