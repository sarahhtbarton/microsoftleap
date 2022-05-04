package internal

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarahhtbarton/microsoftleap/internal/matrix"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	rows, columns, err := helper.ConvertMaptoInts(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	matrix, err := matrix.GenerateFibMatrix(rows, columns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matrix)
}
