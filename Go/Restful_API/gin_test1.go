package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
func PostHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
func posts(c *gin.Context) {
	body := c.Request.Body
	val, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	c.JSON(200, gin.H{
		"posts": string(val),
	})
}

func postwithquery(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   c.Query("id"),
		"name": c.Query("name"),
	})
}

func main() {
	fmt.Println("Hello")
	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/posthomepage", PostHomePage)
	r.POST("/posts", posts)
	r.GET("/postwithquery", postwithquery)
	r.Run()

}
