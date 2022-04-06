package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarahhtbarton/microsoftleap/internal"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", internal.Handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}