package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sarahhtbarton/microsoftleap/internal/helper"
)

//TODO: do i need to break up this function more? "Single responsibility principle"
func Handler(w http.ResponseWriter, r *http.Request) {

	// Get the route variables
	vars := mux.Vars(r)
	numRows := vars["rows"]
	numColumns := vars["columns"]

	//Convert map values from string to int
	rows, _ := strconv.Atoi((numRows))
	columns, _ := strconv.Atoi((numColumns))

	//TODO: error handling 
	//How do you actually do errors? `return error.new("needs to be positive num")` ?
	if len(numRows) == 0 || len(numColumns) == 0 {
		fmt.Println("Neither columns nor rows can be empty")
	}
	if rows < 1 || rows > 92 || columns < 1 || columns > 92{
		fmt.Println("You must enter a number greater than 0 and less than 92")
	}
	if rows * columns > 92 {
		fmt.Println("Can only support a matrix with maximum of 92 numbers")
	}

	//Get the matrix for the given parameters
	matrix := helper.FibonacciMatrix(rows, columns)

	//TODO: Learn more about what these different things do
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	//Encode the matrix to JSON
	res, err := json.Marshal(matrix)
	if err != nil {
		//TODO: which is the right way to handle an error?
		fmt.Println("error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//Update response (Which of the two options below do I use?)
	json.NewEncoder(w).Encode(res) //What would this do?
	// w.Write(res)
}
