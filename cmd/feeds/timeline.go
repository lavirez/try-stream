package feeds

import (
    "github.com/GetStream/stream-go2/v8"
)

type TimeLine struct{ 
    FeedParent
    Owner string `json:"owner"`
}

/* 
    each user has unique timeline
*/

func (time_line *TimeLine) createTimeLine(client *stream.Client, activity *stream.Activity) (*stream.FlatFeed, error) {
    streamTimeLine, err := client.FlatFeed("timeline", time_line.Owner)
    if err != nil { 
        return nil, err
    }
    println(streamTimeLine)
    return streamTimeLine, nil
}

func (time_line *TimeLine) getTimeLine(client *stream.Client) { 
}

func (time_line *TimeLine) updateTimeLine(client *stream.Client) {
}

func (time_line *TimeLine) deleteTimeLine(client *stream.Client) { 
}

// reciever to follow some entity
func (time_line *TimeLine) follow(feed FeedParent) { 
}
