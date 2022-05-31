package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sarahhtbarton/microsoftleap/internal/matrix"
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

	r.HandleFunc("/", handlerIndex).Methods("GET")
	
	r.HandleFunc("/matrix", handlerMatrix).Methods("GET")

	r.Path("/metrics").Handler(promhttp.Handler())

	return r
}

func ParseMatrixDimensions(vars map[string][]string) (int, int, error) {

	if len(vars["rows"]) == 0 || len(vars["columns"]) == 0 {
		return 0, 0, errors.New("Please enter both row and column parameters")
	}

	rowsStr := vars["rows"][0]
	columnsStr := vars["columns"][0]

	rows, err := strconv.Atoi(rowsStr)
	if err != nil {
		return 0, 0, errors.New("Please enter an *integer* for your desired number of rows")
	}
	
	columns, err := strconv.Atoi(columnsStr)
	if err != nil {
		return 0, 0, errors.New("Please enter an *integer* for your desired number of columns")
	}

	return rows, columns, nil
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome")
}

func handlerMatrix(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	rows, columns, err := ParseMatrixDimensions(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	matrix, err := matrix.GenerateFibMatrix(rows, columns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matrix)
}
