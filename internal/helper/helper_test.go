package helper

import (
	"errors"
	"testing"
)

func TestConvertMaptoInts(t *testing.T) {

	data := []struct {
		name    string
		vars    map[string]string
		rows    int
		columns int
		err     error
	}{
		{
			name:    "Row Invalid",
			vars:    map[string]string{"rows": "d", "columns": "2"},
			rows:    0,
			columns: 0,
			err:     errors.New("Please enter an *integer* for your desired number of rows"),
		},
		{
			name:    "Column Invalid",
			vars:    map[string]string{"rows": "7", "columns": "f"},
			rows:    0,
			columns: 0,
			err:     errors.New("Please enter an *integer* for your desired number of columns"),
		},
		{
			name:    "All Valid",
			vars:    map[string]string{"rows": "3", "columns": "5"},
			rows:    3,
			columns: 5,
			err:     nil,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			rows, columns, err := ConvertMaptoInts(d.vars)
			if rows != d.rows {
				t.Errorf("Expected %d, got %d", d.rows, rows)
			}
			if columns != d.columns {
				t.Errorf("Expected %d, got %d", d.columns, columns)
			}
			if err == nil && d.err != nil {
				t.Errorf("Expected %d, got %d", d.err, err)
			} else if err != nil && d.err == nil {
				t.Errorf("Expected %d, got %d", d.err, err)
			} else if err != nil && d.err != nil {
				if err.Error() != d.err.Error() {
					t.Errorf("Expected %d, got %d", d.err, err)
				}
			}
		})
	}
}

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
			m:        3,
			n:        4,
			expected: [][]int64{{0, 1, 1, 2}, {34, 55, 89, 3}, {21, 13, 8, 5}},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := FibonacciMatrix(d.m, d.n)
			if len(result) != len(d.expected) {
				t.Errorf("Expeced %d, got %d", d.expected, result)
			}
			//TODO: compare column length as well
			//TODO: compare each element
			if result[d.m-1][d.n-1] != d.expected[d.m-1][d.n-1] {
				t.Errorf("Expeced %d, got %d", d.expected, result)
			}
		})
	}
}
