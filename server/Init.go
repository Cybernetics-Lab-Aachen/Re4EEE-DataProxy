package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Cybernetics-Lab-Aachen/Re4EEE-DataProxy/twitter"
	"github.com/Cybernetics-Lab-Aachen/Re4EEE-DataProxy/wikipedia"
)

func init() {

	serverMUX := http.NewServeMux()
	serverMUX.HandleFunc("/twitter", twitter.HandlerGetTweets)
	serverMUX.HandleFunc("/wikipedia/getArticleByTitle", wikipedia.HandlerGetArticle)

	server = &http.Server{}
	server.Addr = os.Getenv("Re4EEEDataProxy_ServerIfacePort")
	server.Handler = serverMUX
	server.SetKeepAlivesEnabled(false)
	server.ReadTimeout = 30 * time.Minute
	server.WriteTimeout = 30 * time.Minute
	server.ErrorLog = log.New(os.Stdout, ``, log.LstdFlags)

	log.Printf("Server is configured on '%s'", os.Getenv("Re4EEEDataProxy_ServerIfacePort"))
}
