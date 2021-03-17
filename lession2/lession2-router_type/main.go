package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// GET 获取所有的文章信息
	r.GET("/posts",func(c *gin.Context){
		c.String(http.StatusOK,"GET")
	})

	r.POST("/posts",func(c *gin.Context){
		c.String(http.StatusOK,"POST")
	})

	r.PUT("/posts/:id",func(c *gin.Context){
		c.String(http.StatusOK,fmt.Sprintf("PUT id:%s",c.Param("id")))
	})

	r.DELETE("/posts",func(c *gin.Context){
		c.String(http.StatusOK,"DELETE")
	})

	//匹配所有请求方法
	r.Any("users", func(c *gin.Context) {
		c.String(200,"any")
	})

	r.Run(":3005")
}
