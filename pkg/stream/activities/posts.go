package activities

import (
	"context"
    "encoding/json"

	"github.com/GetStream/stream-go2/v8"
)

// takes in ID and Owner to create a stream related post
type Post struct{ 
    ID      string `json:"id"`
    Owner   string `json:"owner"`
    Content string `json:"content"`
}

func (post *Post) CreateStreamPost(client *stream.Client, feed *stream.FlatFeed) []byte { 
    streamPost := stream.Activity { 
        Actor: post.Owner,
        Verb: "post",
        Object: "THis is the Object",
        ForeignID: post.ID,
    }

    response, err := feed.AddActivity(context.TODO(), streamPost)
    if err != nil { 
        return nil
    }
    marshalled_response, err := json.Marshal(response)
    if err != nil { 
        return nil
    }
    return marshalled_response
}
