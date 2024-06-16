package main

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var mutex = &sync.Mutex{}
var counter = 0

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// read counter route
	r.GET("/counter", func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		c.JSON(200, gin.H{
			"counter": counter,
		})
	})

	// increment counter route
	r.GET("/hit", func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		counter++
		c.JSON(200, gin.H{
			"counter": counter,
		})
	})

	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
