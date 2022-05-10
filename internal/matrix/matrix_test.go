package matrix

import (
	"errors"
	"testing"
)

func TestClientErrorHandling(t *testing.T) {

	data := []struct {
		name     string
		rows     int
		columns  int
		expected [][]int64
		err		 error
	}{
		{
			name:     "Invalid Range",
			rows:     0,
			columns:  7,
			expected: nil,
			err: errors.New("You must enter a number greater than 0 and less than 93"),
		},
		{
			name:     "Invalid Total Size",
			rows:     10,
			columns:  11,
			expected: nil,
			err: errors.New("Can only support a matrix with maximum of 92 numbers"),
		},
		{
			name:     "Valid Matrix",
			rows:     3,
			columns:  4,
			expected: [][]int64{{0, 1, 1, 2}, {34, 55, 89, 3}, {21, 13, 8, 5}},
			err:	  nil,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := GenerateFibMatrix(d.rows, d.columns)
			
			// Compares errors
			if err == nil && d.err != nil {
				t.Errorf("Expected %d, got %d", d.err, err)
			} else if err != nil && d.err == nil {
				t.Errorf("Expected %d, got %d", d.err, err)
			} else if err != nil && d.err != nil {
				if err.Error() != d.err.Error() {
					t.Errorf("Expected %d, got %d", d.err, err)
				}
			}
			
			// Compares column length (num of slices)
			if len(result) != len(d.expected) {
				t.Errorf("Expeced %d, got %d", d.expected, result)
			}
			//Compares row length (num of elements in each slice) 
			for _, slice := range d.expected {
				for _, resultSlice := range result {
					if len(slice) != len(resultSlice) {
						t.Errorf("Expeced %d, got %d", len(slice), len(resultSlice))
					}
				}
			}
			//Compares each element
			for i, slice := range d.expected {
				for j, element := range slice {
					if element != result[i][j] {
						t.Errorf("Expeced %d, got %d", element, result[i][j])
					}
				}
			}
		})
	}
}