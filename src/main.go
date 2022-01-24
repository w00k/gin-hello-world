package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func main() {
	fmt.Println("gin-hello-worl start")
	router := gin.Default()
	router.GET("/hello-world", helloWorld)
	router.Run()
}
