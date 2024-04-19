package main

import (
	"github.com/CyberFatherRT/gorl"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.GET("/api/v1/get_link", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"link": RandStringRunes(5),
		})
	})

	r.Run(":8000")
}
