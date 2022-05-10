package server

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMatrixDimensions(t *testing.T) {

	data := []struct {
		name    string
		vars    map[string][]string
		rows    int
		columns int
		err     error
	}{
		{
			name:    "Empty Map",
			vars:    map[string][]string{},
			rows:    0,
			columns: 0,
			err:     errors.New("Please enter both row and column parameters"),
		},
		{
			name:    "Nil Map",
			vars:    nil,
			rows:    0,
			columns: 0,
			err:     errors.New("Please enter both row and column parameters"),
		},
		{
			name:    "Missing Row Parameter",
			vars:    map[string][]string{"columns": {"2"}},
			rows:    0,
			columns: 0,
			err:     errors.New("Please enter both row and column parameters"),
		},
		{
			name:    "Missing Column Parameter",
			vars:    map[string][]string{"rows": {"d"}},
			rows:    0,
			columns: 0,
			err:     errors.New("Please enter both row and column parameters"),
		},
		{
			name:    "Row Invalid",
			vars:    map[string][]string{"rows": {"d"}, "columns": {"2"}},
			rows:    0,
			columns: 0,
			err:     errors.New("Please enter an *integer* for your desired number of rows"),
		},
		{
			name:    "Column Invalid",
			vars:    map[string][]string{"rows": {"7"}, "columns": {"f"}},
			rows:    0,
			columns: 0,
			err:     errors.New("Please enter an *integer* for your desired number of columns"),
		},
		{
			name:    "All Valid",
			vars:    map[string][]string{"rows": {"3"}, "columns": {"5"}},
			rows:    3,
			columns: 5,
			err:     nil,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			rows, columns, err := ParseMatrixDimensions(d.vars)

			require.Equal(t, d.rows, rows, "Mismatched rows")

			require.Equal(t, d.columns, columns, "Mismatched columns")

			require.Equal(t, d.err, err, "Mismatched error messages")
		})
	}
}