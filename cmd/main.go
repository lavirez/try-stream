package main

import (
	"github.com/gin-gonic/gin"
)


func main() { 
    r := gin.Default()
    r.GET("/", func(c *gin.Context) { 
        c.JSON(200, gin.H{ 
            "message": "API is up and running!",
        })
    })
    r.POST("/post", CreatePostHandler)
    r.Run()
}
