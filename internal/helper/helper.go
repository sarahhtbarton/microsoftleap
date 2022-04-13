package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func ConvertMaptoInts (vars map[string]string) (int, int, error) {
	
	numRows := vars["rows"]
	numColumns := vars["columns"]

	rows, err := strconv.Atoi(numRows)
	columns, err := strconv.Atoi(numColumns)

	return rows, columns, err

}

func ClientErrorHandling (rows int, columns int) error {

	if rows < 1 || rows > 92 || columns < 1 || columns > 92 {
		return errors.New("You must enter a number greater than 0 and less than 92")
	}
	
	if rows*columns > 92 {
		return errors.New("Can only support a matrix with maximum of 92 numbers")
	}

	return nil
}

func ConvertMatrixToJason (matrix [][]int64) ([]byte, error) {

	res, err := json.Marshal(matrix)
	if err != nil {
		return nil, errors.New("Error converting Matrix to JSON string")
	}

	return res, nil
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
	
   fmt.Println(fibMatrix)
	return fibMatrix
}