package helper

import (
	"errors"
	"strconv"
)

func ConvertMaptoInts (vars map[string]string) (int, int, error) {

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

func ClientErrorHandling (rows int, columns int) error {

	//the max value for a int64 is 9,223,372,036,854,770,000 
	//the largest fibonacci number int64 can handle is the 92nd number
	//this function puts in constraints so that the server never encounters a Fibonacci number that's too large to handle
	if rows < 1 || rows >= 92 || columns < 1 || columns >= 92 {
		return errors.New("You must enter a number greater than 0 and less than 93")
	}
	
	if rows*columns > 92 {
		return errors.New("Can only support a matrix with maximum of 92 numbers")
	}

	return nil
}

func FibonacciMatrix (m int, n int) [][]int64 {
	//m = number of lists (height/rows)
	//n = number of elements in the list (width/columns)

	//Make empty matrix
	fibMatrix := make([][]int64, 0)
	for i := 0; i < m; i++ {
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
			n0 := n1 + n2
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
	return fibMatrix
}