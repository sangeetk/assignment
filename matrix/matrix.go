package matrix

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
)

type Matrix struct {
	m [][]int
}

func New(r *http.Request) (*Matrix, error) {

	// Parse CSV from the request body.
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	var m Matrix
	m.m = [][]int{}

	// Validate if the CSV is a valid square matrix of integers.
	for _, row := range records {
		if len(row) != len(records) {
			return nil, fmt.Errorf("not a square matrix")
		}

		vals := []int{}
		for _, v := range row {
			val, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("convertion error '%v'", err)
			}
			vals = append(vals, val)
		}
		m.m = append(m.m, vals)
	}

	// Return matrix
	return &m, nil
}

func (m *Matrix) Invert() *Matrix {

	for i := 0; i < len(m.m); i++ {
		for j := i + 1; j < len(m.m[i]); j++ {
			t := m.m[i][j]
			m.m[i][j] = m.m[j][i]
			m.m[j][i] = t
		}
	}

	return m
}

func (m *Matrix) Sum() int {
	sum := 0

	for _, row := range m.m {
		for _, v := range row {
			sum += v
		}
	}

	return sum
}

func (m *Matrix) Multiply() int {
	product := 1

	for _, row := range m.m {
		for _, v := range row {
			product *= v
		}
	}

	return product
}

func (m *Matrix) String() string {
	var str string

	for i, row := range m.m {
		for j, val := range row {
			switch {
			case j < len(row)-1:
				str += fmt.Sprintf("%d,", val)
			default:
				str += fmt.Sprintf("%d", val)
			}
		}
		if i < len(m.m)-1 {
			str += "\n"
		}
	}

	return str
}

func (m *Matrix) Flatten() string {
	var str string

	for i, row := range m.m {
		for j, val := range row {
			switch {
			case i == len(m.m)-1 && j == len(row)-1:
				str += fmt.Sprintf("%d", val)
			default:
				str += fmt.Sprintf("%d,", val)
			}
		}
	}

	return str
}
