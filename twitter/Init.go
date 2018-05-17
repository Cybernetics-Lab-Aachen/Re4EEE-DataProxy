package twitter

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
)

func init() {
	log.Println("Init the Twitter package.")

	databaseHostname := os.Getenv("Re4EEEDataProxy_TwitterDBHostname")
	databaseDB = os.Getenv("Re4EEEDataProxy_TwitterDBDatabaseName")
	databaseUsername = os.Getenv("Re4EEEDataProxy_TwitterDBUsername")
	databasePassword = os.Getenv("Re4EEEDataProxy_TwitterDBPassword")
	databaseCollectionName = os.Getenv("Re4EEEDataProxy_TwitterDBCollectionName")

	// Connect to MongoDB:
	if newSession, errDial := mgo.Dial(databaseHostname); errDial != nil {
		log.Printf("Was not able to connect to the Twitter database: %s.", errDial)
		os.Exit(101)
	} else {
		mainSession = newSession
	}

	// Use the correct database:
	db := mainSession.DB(databaseDB)
	if db == nil {
		log.Printf("Was not able to connect to the Twitter database.")
		os.Exit(102)
	}

	// Login:
	if errLogin := db.Login(databaseUsername, databasePassword); errLogin != nil {
		log.Printf("Was not able to connect to the Twitter database: %s.", errLogin)
		os.Exit(103)
	}

	// In case of write operations, wait for the majority of servers to be done:
	mainSession.SetSafe(&mgo.Safe{WMode: "majority"})

	// Set the consistency mode to read from any secondary server and write to the primary.
	// Copied sessions can overwrite this setting of necessary.
	mainSession.SetMode(mgo.Eventual, true)

	log.Println("Init the Twitter package was successfully.")
}
