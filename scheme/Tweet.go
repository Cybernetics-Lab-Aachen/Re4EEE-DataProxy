package scheme

import "time"

type Tweet struct {
	TextEN       string    `bson:"TextEN" json:"Text"`
	TweetTimeUTC time.Time `bson:"TweetTimeUTC" json:"TimeUTC"`
}

type TwitterResults struct {
	Tweets []Tweet `json:"Tweets"`
}
