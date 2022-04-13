package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarahhtbarton/microsoftleap/internal/helper"
)

func Handler(w http.ResponseWriter, r *http.Request) {

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

	res, err3 := helper.ConvertMatrixToJason(matrix)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}

	//TODO: Learn more about what these different things do
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// json.NewEncoder(w).Encode(res) //What would this do?
}
