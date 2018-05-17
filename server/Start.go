package server

import "log"

func Start() {
	log.Print("Start the server now.")
	server.ListenAndServe()
}
