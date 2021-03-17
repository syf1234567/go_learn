package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//加载静态文件
	r.Static("/images", "./images")

	r.StaticFS("/static", http.Dir("./static"))

	//  加载单独的静态文件
	r.StaticFile("index", "index.html")
	r.Run()
}
