package matrix

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrorInvalidIntMatrix = errors.New("invalid int matrix")
	ErrorZeroSizeMatrix   = errors.New("zero size matrix")
	ErrorNonSquareMatrix  = errors.New("not a square matrix")
)

// A square matrix of int.
type Matrix [][]int

// Parses CSV records and creates a square matrix of int values
func Parse(records [][]string) (Matrix, error) {

	var m Matrix

	// Validate & Convert CSV into matrix of integers.
	for _, record := range records {
		row := []int{}

		for _, v := range record {
			n, err := strconv.Atoi(v)
			if err != nil {
				// Invalid (non integer) characters
				return nil, ErrorInvalidIntMatrix
			}
			row = append(row, n)
		}
		m = append(m, row)
	}

	// Check for zero size matrix
	if len(m) == 0 {
		return nil, ErrorZeroSizeMatrix
	}

	// Check if number of values in each row is same as total number of rows.
	for _, row := range m {
		if len(row) != len(m) {
			return nil, ErrorNonSquareMatrix
		}
	}

	// Return int matrix
	return m, nil
}

// Invert the matrix by swapping upper & lower triangular values
func (m Matrix) Invert() Matrix {

	for i := 0; i < len(m); i++ {
		for j := i + 1; j < len(m[i]); j++ {
			t := m[i][j]
			m[i][j] = m[j][i]
			m[j][i] = t
		}
	}

	return m
}

// Returns the sum of all values in a matrix
func (m Matrix) Sum() int {
	sum := 0

	for _, row := range m {
		for _, n := range row {
			sum += n
		}
	}

	return sum
}

// Returns the product of all values in a matrix
func (m Matrix) Multiply() int {
	product := 1

	for _, row := range m {
		for _, n := range row {
			product *= n
		}
	}

	return product
}

// Return the matrix as string in 2D matrix form using csv
func (m Matrix) String() string {
	var str string

	for i, row := range m {
		for j, n := range row {
			switch {
			case j < len(row)-1:
				str += fmt.Sprintf("%d,", n)
			default:
				// Don't add ',' after last value in the row
				str += fmt.Sprintf("%d", n)
			}
		}
		if i < len(m)-1 {
			str += "\n"
		}
	}

	return str
}

// Return the matrix in flattened form as csv string
func (m Matrix) Flatten() string {
	var str string

	for i, row := range m {
		for j, n := range row {
			switch {
			case i == len(m)-1 && j == len(row)-1:
				str += fmt.Sprintf("%d", n)
			default:
				str += fmt.Sprintf("%d,", n)
			}
		}
	}

	return str
}
