package server

import "log"

func Start() {
	go func() {
		log.Print("Start the server now.")
		server.ListenAndServe()
	}()
}
