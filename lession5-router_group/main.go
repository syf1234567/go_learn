package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//r.GET("/posts",GetHandler)
	//r.POST("/posts",PostHandler)
	//r.DELETE("/posts/1",DeleteHandler)

	//p := r.Group("/posts")
	//{
	//	p.GET("",GetHandler)
	//	p.POST("",PostHandler)
	//	p.DELETE("/:id",DeleteHandler)
	//}

	//v1 := r.Group("/api/v1")
	//{
	//	v1.POST("/login", loginEndpoint)
	//	v1.POST("/submit", submitEndpoint)
	//	v1.POST("/read", readEndPoint)
	//}

	//v2 := r.Group("/v2")
	//{
	//	v2.POST("/login", loginEndpoint)
	//	v2.POST("/submit", submitEndpoint)
	//	v2.POST("/read", readEndPoint)
	//}

	api := r.Group("/api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("/posts", GetHandler)
			v1.POST("/posts", PostHandler)
			v1.DELETE("/:id", DeleteHandler)
		}
	}

	r.Run()
}

func readEndPoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "readEndpoint",
	})
}

func submitEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "submitEndpoint",
	})
}
func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "loginEndpoint",
	})
}

func DeleteHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DELETE",
	})
}

func PostHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post",
	})
}

func GetHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get",
	})
}
