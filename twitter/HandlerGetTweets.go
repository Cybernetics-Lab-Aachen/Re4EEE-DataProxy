package twitter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Cybernetics-Lab-Aachen/Re4EEE-DataProxy/scheme"
	"github.com/globalsign/mgo/bson"
)

// HandlerGetTweets provides the functionality to retrieve Tweets out of the database.
func HandlerGetTweets(response http.ResponseWriter, request *http.Request) {

	tweets := make([]scheme.Tweet, 0)
	yesterday := time.Now().UTC().AddDate(0, 0, -1)
	fromDate := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, time.UTC)
	toDate := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 59, 0, time.UTC)

	session := mainSession.Copy()
	database := session.DB(databaseDB)
	database.Login(databaseUsername, databasePassword)
	defer session.Close()

	database.C(databaseCollectionName).Find(bson.M{
		"TweetTimeUTC": bson.M{
			"$gt": fromDate,
			"$lt": toDate,
		},
	}).Select(bson.M{"TextEN": 1, "TweetTimeUTC": 1}).Sort("-TweetTimeUTC").Limit(1000000).All(&tweets)

	result := scheme.TwitterResults{
		Tweets: tweets,
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Fprintln(response, string(jsonBytes))
}
