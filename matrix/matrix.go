package matrix

import (
	"fmt"
	"strconv"
)

// A square matrix of int.
type Matrix [][]int

// Checks if the matrix square
func isSquare(records [][]string) bool {

	// Check if number of values in each row is same as total number of rows.
	for _, record := range records {
		if len(record) != len(records) {
			return false
		}
	}
	return true
}

// Parses CSV records and creates a square matrix of int values
func Parse(records [][]string) (*Matrix, error) {

	var m Matrix

	if !isSquare(records) {
		return nil, fmt.Errorf("not a square matrix")
	}

	// Validate & Convert CSV into matrix of integers.
	for _, record := range records {
		row := []int{}

		for _, v := range record {
			n, err := strconv.Atoi(v)
			if err != nil {
				// Invalid (non integer) characters
				return nil, fmt.Errorf("convertion error '%v'", err)
			}
			row = append(row, n)
		}
		m = append(m, row)
	}

	// Return int matrix
	return &m, nil
}

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

func (m Matrix) Sum() int {
	sum := 0

	for _, row := range m {
		for _, n := range row {
			sum += n
		}
	}

	return sum
}

func (m Matrix) Multiply() int {
	product := 1

	for _, row := range m {
		for _, n := range row {
			product *= n
		}
	}

	return product
}

func (m Matrix) String() string {
	var str string

	for i, row := range m {
		for j, n := range row {
			switch {
			case j < len(row)-1:
				str += fmt.Sprintf("%d,", n)
			default:
				str += fmt.Sprintf("%d", n)
			}
		}
		if i < len(m)-1 {
			str += "\n"
		}
	}

	return str
}

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
