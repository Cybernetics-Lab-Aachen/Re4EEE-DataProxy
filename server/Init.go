package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Cybernetics-Lab-Aachen/Re4EEE-DataProxy/twitter"
)

func init() {

	serverMUX := http.NewServeMux()
	serverMUX.HandleFunc("/twitter", twitter.HandlerGetTweets)

	server = &http.Server{}
	server.Addr = os.Getenv("Re4EEEDataProxy_ServerIfacePort")
	server.Handler = serverMUX
	server.SetKeepAlivesEnabled(false)
	server.ReadTimeout = 3 * time.Second
	server.WriteTimeout = 30 * time.Minute
	server.ErrorLog = log.New(os.Stdout, ``, log.LstdFlags)
}
