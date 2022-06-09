package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/randomnumber", handlerRandom).Methods("GET")

	value, ok := os.LookupEnv("PORT")
	if ok {
		http.ListenAndServe(value, r)
	} else {
		http.ListenAndServe(":8080", r)
	}
}

func handlerRandom(w http.ResponseWriter, r *http.Request) {

	fmt.Println(os.Getenv("MYVAR"))

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 92
	randomNumber := rand.Intn(max-min+1) + min

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(randomNumber)
}
