package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarahhtbarton/microsoftleap/internal"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/rows/{rows}/columns/{columns}", internal.Handler)
	log.Fatal(http.ListenAndServe(":8080", r))

	//Test:
	//http://localhost:8080/rows/3/columns/4
}