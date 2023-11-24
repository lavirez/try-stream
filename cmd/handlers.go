package main

import (
	"net/http"

	streamclient "github.com/alire-alavi/try-stream/internal/stream_client"
	// "github.com/GetStream/stream-go2/v8"
	"github.com/alire-alavi/try-stream/cmd/activities"
	"github.com/gin-gonic/gin"
)

type CreatePostInput struct { 
    ID      string `json:"id"`
    Content string `json:"content"`
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
    response := post.CreateStreamPost(client, feed)
    c.JSON(http.StatusOK, gin.H{"data": response})
}
