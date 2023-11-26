package feeds

type FeedParent struct {
    ID       string `json:"id"`
    FeedType string `json:"feed_type"`
    FeedGroup string `json:"feed_group"`
}

const ( 
    NOTIFICATION = "notification"
    USER = "user"
    TIMELINEAGGREGATED = "timeline_aggregated"
    TIMELINE = "timeline"
)
