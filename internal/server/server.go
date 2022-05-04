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

	rows, columns, err1 := helper.ConvertMaptoInts(vars)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusNotAcceptable)
		return
	}
	
	err2 := helper.ClientErrorHandling(rows, columns)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusRequestedRangeNotSatisfiable)
		return
	}

	matrix := helper.FibonacciMatrix(rows, columns)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matrix)
}
