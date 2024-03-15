package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var db map[string]string = make(map[string]string)
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type AddLink struct {
	link string
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func StringWithCharset() string {
	b := make([]byte, 13)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func add_link(c *gin.Context) {

	var body AddLink

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "ERROR: cannot parse json")
		return
	}

	rand_string := StringWithCharset()

	db[rand_string] = body.link

	c.JSON(http.StatusOK, gin.H{"short_link": rand_string})
}

func transfer(c *gin.Context) {
	short_link := c.Param("link")

	link, err := db[short_link]

	if err == false {
		c.Redirect(http.StatusNotFound, "/")
	}

	c.Redirect(http.StatusMovedPermanently, link)
	c.Abort()
}

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World")
	})
	r.GET("/:link", transfer)
	r.GET("/ping", pong)
	r.POST("/add_link", add_link)

	r.Run(":8080")
}
