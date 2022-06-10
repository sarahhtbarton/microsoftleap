package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/randomnumber", handlerRandom).Methods("GET")

	// set env variable and run with: `PORT=":8081" go run main.go``
	value, ok := os.LookupEnv("PORT")
	if ok {
		http.ListenAndServe(value, r)
	} else {
		http.ListenAndServe(":8080", r)
	}
}

func handlerRandom(w http.ResponseWriter, r *http.Request) {

	// TODO: return a map with {"row:" x, "column": y}
	// TODO: throw it in for loop to make sure x*y is not greater than 92
	
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 92
	randomNumber := rand.Intn(max-min+1) + min

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(randomNumber)
}
