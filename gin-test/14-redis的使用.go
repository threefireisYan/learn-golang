package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//初始化基于redis的引擎
	//参数说明
	//第 1 个参数 - redis最大的空闲连接数
	//第 2 个参数 - 通信协议tcp或udp
	//第 3 个参数 - redis地址，格式，host:port
	//第 4 个参数 - redis密码
	//第 5 个参数 - session加密秘钥
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8094")
}
