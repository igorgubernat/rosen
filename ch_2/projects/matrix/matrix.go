package matrix

import (
	"fmt"
)

// Matrix represents mathematical matrix
type Matrix struct {
	els     [][]float64
	numRows int
	numCols int
}

// NewMatrix is constructor for Matrix
func NewMatrix(els [][]float64) (*Matrix, error) {
	numCols := len(els[0])
	for _, row := range els {
		if len(row) != numCols {
			return nil, fmt.Errorf("%v is not valid matrix", els)
		}
	}

	return &Matrix{
		els:     els,
		numRows: len(els),
		numCols: len(els[0]),
	}, nil
}

func (m *Matrix) col(j int) []float64 {
	var c []float64
	for _, el := range m.els {
		c = append(c, el[j])
	}

	return c
}

// X is matrix multiplication
func (m *Matrix) X(other *Matrix) (*Matrix, error) {
	if m.numCols != other.numRows {
		return nil, fmt.Errorf("First matrix has %d columns which is not equal to second matrix' number of rows %d", m.numCols, other.numRows)
	}

	els := make([][]float64, m.numRows)
	for i := range els {
		els[i] = make([]float64, other.numCols)
	}

	for i, row := range m.els {
		for j := 0; j < other.numCols; j++ {
			col := other.col(j)
			var res float64
			for k := range row {
				res += row[k] * col[k]
			}
			els[i][j] = res
		}
	}

	return NewMatrix(els)
}
