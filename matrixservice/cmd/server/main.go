package main

import (
	"log"

	"github.com/sarahhtbarton/microsoftleap/matrixservice/internal/server"
)

func main() {

	srvr := server.NewServer(":8080")
	
	log.Fatal(srvr.ListenAndServe())
}