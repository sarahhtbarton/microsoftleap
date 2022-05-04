package main

import (
	"log"

	"github.com/sarahhtbarton/microsoftleap/internal/server"
)

func main() {

	srvr := server.NewServer(":8080")
	
	log.Fatal(srvr.ListenAndServe())
}