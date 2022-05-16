package matrix

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateFibMatrix(t *testing.T) {

	data := []struct {
		name     string
		rows     int
		columns  int
		expected [][]int64
		err      error
	}{
		{
			name:     "Invalid Range",
			rows:     0,
			columns:  7,
			expected: nil,
			err:      errors.New("You must enter a number greater than 0 and less than 93"),
		},
		{
			name:     "Invalid Total Size",
			rows:     10,
			columns:  11,
			expected: nil,
			err:      errors.New("Can only support a matrix with maximum of 92 numbers"),
		},
		{
			name:     "Valid Matrix",
			rows:     3,
			columns:  4,
			expected: [][]int64{{0, 1, 1, 2}, {34, 55, 89, 3}, {21, 13, 8, 5}},
			err:      nil,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			expected, err := GenerateFibMatrix(d.rows, d.columns)

			require.Equal(t, d.expected, expected, "Mismatched result")

			require.IsType(t, d.err, err, "Mismatched errors")
		})
	}
}
