package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	fmt.Fprintln(w, "Hello world")
	fmt.Fprintln(w, "Vars dict:", vars)
}
