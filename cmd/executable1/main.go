package main

import (
	"log"
	"net/http"

	"github.com/sarahhtbarton/microsoftleap/internal"
)

func main() {
	http.HandleFunc("/", internal.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}