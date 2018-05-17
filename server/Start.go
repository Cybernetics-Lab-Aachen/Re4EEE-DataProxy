package server

import "log"

func Start() {
	log.Print("Start the server now.")
	errServer := server.ListenAndServe()
	if errServer != nil {
		log.Printf("Error while running the server: %s", errServer)
	}
}
