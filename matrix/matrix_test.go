package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMatrix(t *testing.T) {

	type TestCase struct {
		name     string
		records  [][]string
		expected error
	}

	tests := []TestCase{
		{
			name:     "Check for empty matrix",
			records:  [][]string{},
			expected: ErrorZeroSizeMatrix,
		},
		{
			name: "Check for 1x1 matrix",
			records: [][]string{
				{"1"},
			},
			expected: nil,
		},
		{
			name: "Check for 2x2 matrix of float",
			records: [][]string{
				{"1", "2"},
				{"3.2", "4.1"},
			},
			expected: ErrorInvalidIntMatrix,
		},
		{
			name: "Check for mixed int & float values",
			records: [][]string{
				{"1.0", "2"},
				{"4", "-5"},
			},
			expected: ErrorInvalidIntMatrix,
		},
		{
			name: "Check for non numeric characters",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "x", "6"},
				{"7", "8", "9"},
			},
			expected: ErrorInvalidIntMatrix,
		},
		{
			name: "Check for mixed +ve & -ve values",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "-5", "6"},
				{"-7", "8", "+9"},
			},
			expected: nil,
		},
		{
			name: "Check for non-square matrix",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
			},
			expected: ErrorNonSquareMatrix,
		},
	}

	for _, test := range tests {
		_, err := Parse(test.records)
		assert.Equal(t, err, test.expected, test.name)
	}

}

func TestStringConversion(t *testing.T) {

	type TestCase struct {
		name     string
		m        Matrix
		expected string
	}

	tests := []TestCase{
		{
			name: "Test 1x1 matrix",
			m: Matrix{
				{1},
			},
			expected: "1",
		},
		{
			name: "Test 2x2 matrix",
			m: Matrix{
				{1, 2},
				{3, 4},
			},
			expected: "1,2\n3,4",
		},
		{
			name: "Test 3x3 matrix",
			m: Matrix{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: "1,2,3\n4,5,6\n7,8,9",
		},
	}

	for _, test := range tests {
		output := test.m.String()
		assert.Equal(t, output, test.expected, test.name)
	}

}

func TestInvertConversion(t *testing.T) {

	type TestCase struct {
		name     string
		m        Matrix
		expected string
	}

	tests := []TestCase{
		{
			name: "Test 1x1 matrix",
			m: Matrix{
				{1},
			},
			expected: "1",
		},
		{
			name: "Test 2x2 matrix",
			m: Matrix{
				{1, 2},
				{3, 4},
			},
			expected: "1,3\n2,4",
		},
		{
			name: "Test 3x3 matrix",
			m: Matrix{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: "1,4,7\n2,5,8\n3,6,9",
		},
	}

	for _, test := range tests {
		output := test.m.Invert().String()
		assert.Equal(t, output, test.expected, test.name)
	}

}

func TestFlatten(t *testing.T) {

	type TestCase struct {
		name     string
		m        Matrix
		expected string
	}

	tests := []TestCase{
		{
			name: "Test 1x1 matrix",
			m: Matrix{
				{1},
			},
			expected: "1",
		},
		{
			name: "Test 2x2 matrix",
			m: Matrix{
				{1, 2},
				{3, 4},
			},
			expected: "1,2,3,4",
		},
		{
			name: "Test 3x3 matrix",
			m: Matrix{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: "1,2,3,4,5,6,7,8,9",
		},
	}

	for _, test := range tests {
		output := test.m.Flatten()
		assert.Equal(t, output, test.expected, test.name)
	}

}

func TestMatrixSum(t *testing.T) {

	type TestCase struct {
		name     string
		m        Matrix
		expected int
	}

	tests := []TestCase{
		{
			name: "Test 1x1 matrix",
			m: Matrix{
				{1},
			},
			expected: 1,
		},
		{
			name: "Test 2x2 matrix",
			m: Matrix{
				{1, 2},
				{3, 4},
			},
			expected: 10,
		},
		{
			name: "Test 2x2 matrix with -ve values",
			m: Matrix{
				{1, -2},
				{-3, 4},
			},
			expected: 0,
		},
		{
			name: "Test 3x3 matrix",
			m: Matrix{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: 45,
		},
	}

	for _, test := range tests {
		output := test.m.Sum()
		assert.Equal(t, output, test.expected, test.name)
	}

}

func TestMatrixMultiply(t *testing.T) {

	type TestCase struct {
		name     string
		m        Matrix
		expected int
	}

	tests := []TestCase{
		{
			name: "Test 1x1 matrix",
			m: Matrix{
				{1},
			},
			expected: 1,
		},
		{
			name: "Test 2x2 matrix",
			m: Matrix{
				{1, 2},
				{3, 4},
			},
			expected: 24,
		},
		{
			name: "Test 2x2 matrix with -ve values",
			m: Matrix{
				{1, -2},
				{3, 4},
			},
			expected: -24,
		},
		{
			name: "Test 3x3 matrix",
			m: Matrix{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: 362880,
		},
		{
			name: "Test 3x3 matrix with 0",
			m: Matrix{
				{1, 2, 3},
				{4, 5, 0},
				{7, 8, 9},
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		output := test.m.Multiply()
		assert.Equal(t, output, test.expected, test.name)
	}

}
