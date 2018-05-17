package twitter

import "github.com/globalsign/mgo"

var (
	mainSession            *mgo.Session = nil // The session for the customer database
	databaseUsername       string       = ``  // The user's name
	databasePassword       string       = ``  // The user's password
	databaseDB             string       = ``  // The database
	databaseCollectionName string       = ``  // The name of the Twitter collection
)
