package main

import (
	"net/http"

	streamclient "github.com/alire-alavi/try-stream/internal/stream_client"
	// "github.com/GetStream/stream-go2/v8"
	"github.com/alire-alavi/try-stream/pkg/stream/activities"
	"github.com/gin-gonic/gin"
)

type CreatePostInput struct { 
    ID      string `json:"id"`
    Content string `json:"content"`
}

type GetPostQuery struct {
	ID   string `uri:"id" binding:"required,uuid"`
}


var client = streamclient.SetUp()

func CreatePostHandler(c *gin.Context) { 
    input := CreatePostInput{}
    if err := c.ShouldBind(&input); err != nil { 
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
   
    post := activities.Post{ 
        ID: input.ID,
        Owner: "666",
        Content: input.Content,
    }
    feed, err := client.FlatFeed("user", "666")
    if err != nil { 
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    bResponse := post.CreateStreamPost(client, feed)
    println(string(bResponse))
    c.JSON(http.StatusOK, gin.H{"data": string(bResponse)})
}

func GetPost(c *gin.Context) { 
}
