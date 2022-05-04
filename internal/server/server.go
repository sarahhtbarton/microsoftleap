package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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

func parseMatrixDimensions(vars map[string]string) (int, int, error) {

	numRows := vars["rows"]
	numColumns := vars["columns"]

	rows, err := strconv.Atoi(numRows)
	if err != nil {
		return 0, 0, errors.New("Please enter an *integer* for your desired number of rows")
	}
	
	columns, err := strconv.Atoi(numColumns)
	if err != nil {
		return 0, 0, errors.New("Please enter an *integer* for your desired number of columns")
	}

	return rows, columns, nil
}

func handler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	rows, columns, err := parseMatrixDimensions(vars)
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
