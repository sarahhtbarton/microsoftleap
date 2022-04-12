package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r) //I think this is used to get dynamic segments from url. Like /page/{page}, you'd use vars["page"] to get the actual page passed

	fmt.Fprintln(w, "Hello world")
	fmt.Fprintln(w, "Vars dict:", vars)

	fibonacciMatrix(3,4) //[[0, 1, 1, 2], [34, 55, 89, 3], [21, 13, 8, 5]]
	fibonacciMatrix(3,2)
	fibonacciMatrix(3,1)
	fibonacciMatrix(0,0)
	fibonacciMatrix(1,1)
	fibonacciMatrix(1,4)
}

//How do you make the return output json? You dont do that here. 
//"Single responsibility principle" -- not the job of this function to do the encoding.
func fibonacciMatrix (m int, n int) [][]int64 {

	//the max value for a int64 is  9,223,372,036,854,770,000 
	//the largest fibonacci number int64 can handle is the 92nd number
	//the 92nd fibonacci number is  7,540,113,804,746,340,000 
	// TODO: validation / error handling. Return both the slice of slices, AND an error. If err != nil ...
		//if n*m > 92... return error


	//Make empty matrix
	fibMatrix := make([][]int64, 0) //TODO: instead of 0 -- can specify number of rows
	for i := 0; i < m; i++ { //TODO: `for i, v := fibMatrix` or `for _, v := fibMatrix`  (this is like enumeration)
		new := make([]int64, n)
		fibMatrix = append(fibMatrix, new)
	}

	k := 0
	l := 0
	var n1 int64 = 1
	var n2 int64 = -1

	for k < m && l < n {
		
		//Handles the first row of remaining rows
		for i := l; i < n; i++ {
			n0 := n1 + n2 //TODO: there is a way to combine these 3 lines with "closures" -- a function that returns a function, where the internal function returned has its own state. Could call newFibGenerator, that returns generateFib, which when called changes its own state of n1, n2 etc. 
			n2 = n1
			n1 = n0
			fibMatrix[k][i] = n0
		}
		k += 1

		//Handles the last column of remaining columns
		for i := k; i < m; i++ {
			n0 := n1 + n2
			n2 = n1
			n1 = n0
			fibMatrix[i][n-1] = n0
		}
		n -= 1

		//Handles the last row of remaining rows
		if k < m {
			for i := n-1; i > l-1; i-- {
				n0 := n1 + n2
				n2 = n1
				n1 = n0
				fibMatrix[m-1][i] = n0
			}
		m -= 1
		}

		//Handles the first column of remaining columns
		if l < n {
			for i:= m-1; i > k-1; i-- {
				n0 := n1 + n2
				n2 = n1
				n1 = n0
				fibMatrix[i][l] = n0
			}
		l += 1
		}
   }
	fmt.Println(fibMatrix)
	return fibMatrix
}
