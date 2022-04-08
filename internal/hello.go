package internal

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	fmt.Fprintln(w, "Hello world")
	fmt.Fprintln(w, "Vars dict:", vars)

	fibonacciMatrix(3,4) //[[0, 1, 1, 2], [34, 55, 89, 3], [21, 13, 8, 5]]
	fibonacciMatrix(3,2)
	fibonacciMatrix(3,1)
	fibonacciMatrix(0,0)
	fibonacciMatrix(1,1)
	fibonacciMatrix(1,4)
}

//TODO: how do you make the return output json?
func fibonacciMatrix (m int, n int) {

	//the max value for a int64 is  9,223,372,036,854,770,000 
	//the largest fibonacci number int64 can handle is the 92nd number
	//the 92nd fibonacci number is  7,540,113,804,746,340,000 



	//Make empty matrix
	//TODO is there a more elegant way to do this? A la `fib_matrix = [[None] * n for i in range(m)]`
	fib_matrix := make([][]int64, 0) //is int64 right here? seems like the type should be a slice....

	for i := 0; i < m; i++ { //is this the way to do `for i in range()`
		new := make([]int64, n)
		fib_matrix = append(fib_matrix, new) //for append, you need to include fib_matrix again?
	}

	k := 0
	l := 0
	var n1 int64 = 1
	var n2 int64 = -1

	for k < m && l < n {
		
		//Handles the first row of remaining rows
		for i := l; i < n; i++ {
			n0 := n1 + n2 //is this DRY??
			n2 = n1
			n1 = n0
			fib_matrix[k][i] = n0
		}
		k += 1

		//Handles the last column of remaining columns
		for i := k; i < m; i++ {
			n0 := n1 + n2
			n2 = n1
			n1 = n0
			fib_matrix[i][n-1] = n0
		}
		n -= 1

		//Handles the last row of remaining rows
		if k < m {
			for i := n-1; i > l-1; i-- {
				n0 := n1 + n2
				n2 = n1
				n1 = n0
				fib_matrix[m-1][i] = n0
			}
		m -= 1
		}

		//Handles the first column of remaining columns
		if l < n {
			for i:= m-1; i > k-1; i-- {
				n0 := n1 + n2
				n2 = n1
				n1 = n0
				fib_matrix[i][l] = n0
			}
		l += 1
		}
   }
	//TODO: return json  
   fmt.Println(fib_matrix)
}
