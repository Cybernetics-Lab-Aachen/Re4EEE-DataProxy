package main

import (
	"log"

	"github.com/Cybernetics-Lab-Aachen/Re4EEE-DataProxy/server"
)

func main() {
	log.Printf("Re4EEE Data Proxy %s", version)
	server.Start()
}
