package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarahhtbarton/microsoftleap/internal/helper"
)

func NewServer(addr string) (*http.Server) {
	
	r := NewHTTPHandler()

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func NewHTTPHandler() (http.Handler) {

	r := mux.NewRouter()

	r.HandleFunc("/rows/{rows}/columns/{columns}", handler).Methods("GET")

	return r
}

func handler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	rows, columns, err := helper.ConvertMaptoInts(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	err = helper.ClientErrorHandling(rows, columns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	matrix := helper.FibonacciMatrix(rows, columns)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matrix)
}
