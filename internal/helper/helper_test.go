package helper

import (
	"errors"
	"testing"
)

func TestClientErrorHandling(t *testing.T) {

	data := []struct {
		name     string
		rows     int
		columns  int
		expected error
	}{
		{
			name:     "Invalid Range",
			rows:     0,
			columns:  7,
			expected: errors.New("You must enter a number greater than 0 and less than 93"),
		},
		{
			name:     "Invalid Total Size",
			rows:     10,
			columns:  11,
			expected: errors.New("Can only support a matrix with maximum of 92 numbers"),
		},
		{
			name:     "Valid Client Inputs",
			rows:     3,
			columns:  7,
			expected: nil,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := ClientErrorHandling(d.rows, d.columns)
			if result == nil && d.expected != nil {
				t.Errorf("Expected %d, got %d", d.expected, result)
			} else if result != nil && d.expected == nil {
				t.Errorf("Expected %d, got %d", d.expected, result)
			} else if result != nil && d.expected != nil {
				if result.Error() != d.expected.Error() {
					t.Errorf("Expected %d, got %d", d.expected, result)
				}
			}
		})
	}
}

func TestFibonacciMatrix(t *testing.T) {

	data := []struct {
		name     string
		m        int
		n        int
		expected [][]int64
	}{
		{
			name:     "Valid Matrix",
			m:        3, // num of lists
			n:        4, // num of elements in list
			expected: [][]int64{{0, 1, 1, 2}, {34, 55, 89, 3}, {21, 13, 8, 5}},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := FibonacciMatrix(d.m, d.n)
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
