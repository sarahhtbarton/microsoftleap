package helper

import "fmt"

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