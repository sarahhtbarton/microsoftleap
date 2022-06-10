package matrix

import (
	"errors"
)

func GenerateFibMatrix(rows int, columns int) ([][]int64, error) {

	//the max value for a int64 is 9,223,372,036,854,770,000 
	//the largest fibonacci number int64 can handle is the 92nd number
	//this function puts in constraints so that the server never encounters a Fibonacci number that's too large to handle
	if rows < 1 || rows >= 92 || columns < 1 || columns >= 92 {
		return nil, errors.New("You must enter a number greater than 0 and less than 93")
	}
	
	if rows*columns > 92 {
		return nil, errors.New("Can only support a matrix with maximum of 92 numbers")
	}

	//Make empty matrix
	fibMatrix := make([][]int64, 0)
	for i := 0; i < rows; i++ {
		new := make([]int64, columns)
		fibMatrix = append(fibMatrix, new)
	}

	k := 0
	l := 0
	var n1 int64 = 1
	var n2 int64 = -1

	for k < rows && l < columns {
		
		//Handles the first row of remaining rows
		for i := l; i < columns; i++ {
			n0 := n1 + n2
			n2 = n1
			n1 = n0
			fibMatrix[k][i] = n0
		}
		k += 1

		//Handles the last column of remaining columns
		for i := k; i < rows; i++ {
			n0 := n1 + n2
			n2 = n1
			n1 = n0
			fibMatrix[i][columns-1] = n0
		}
		columns -= 1

		//Handles the last row of remaining rows
		if k < rows {
			for i := columns-1; i > l-1; i-- {
				n0 := n1 + n2
				n2 = n1
				n1 = n0
				fibMatrix[rows-1][i] = n0
			}
		rows -= 1
		}

		//Handles the first column of remaining columns
		if l < columns {
			for i:= rows-1; i > k-1; i-- {
				n0 := n1 + n2
				n2 = n1
				n1 = n0
				fibMatrix[i][l] = n0
			}
		l += 1
		}
   }
	return fibMatrix, nil
}