package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   int    `uri:"id"`
	Name string `uri:"name"`
}

func main() {
	r := gin.Default()

	//r.GET("/posts", func(c *gin.Context) {
	//	c.String(200, "1")
	//})
	//
	//r.POST("/posts/:id", func(c *gin.Context) {
	//
	//})
	//
	//r.GET("/posts/:id", func(c *gin.Context) {
	//	id := c.Param("id")
	//	c.JSON(200, gin.H{
	//		"id": id,
	//	})
	//})

	r.GET("/:id/:name", func(c *gin.Context) {
		var p Person
		if err := c.ShouldBindUri(&p); err != nil {
			c.Status(404)
			return
		}
		c.JSON(200, gin.H{
			"name": p.Name,
			"id":   p.ID,
		})
	})

	r.Run()
}
