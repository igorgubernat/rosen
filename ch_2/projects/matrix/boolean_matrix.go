package matrix

import (
	"fmt"
)

// BooleanMatrix represents mathematical matrix
type BooleanMatrix struct {
	els     [][]int
	numRows int
	numCols int
}

// NewBooleanMatrix is constructor for BooleanMatrix
func NewBooleanMatrix(els [][]int) (*BooleanMatrix, error) {
	for _, row := range els {
		for _, el := range row {
			if el != 0 && el != 1 {
				return nil, fmt.Errorf("%v is not valid boolean matrix, only 0 or 1 allowed", els)
			}
		}
	}

	numCols := len(els[0])
	for _, row := range els {
		if len(row) != numCols {
			return nil, fmt.Errorf("%v is not valid matrix", els)
		}
	}

	return &BooleanMatrix{
		els:     els,
		numRows: len(els),
		numCols: len(els[0]),
	}, nil
}

func (m *BooleanMatrix) col(j int) []int {
	var c []int
	for _, el := range m.els {
		c = append(c, el[j])
	}

	return c
}

// X is matrix multiplication
func (m *BooleanMatrix) X(other *BooleanMatrix) (*BooleanMatrix, error) {
	if m.numCols != other.numRows {
		return nil, fmt.Errorf("First matrix has %d columns which is not equal to second matrix' number of rows %d", m.numCols, other.numRows)
	}

	els := make([][]int, m.numRows)
	for i := range els {
		els[i] = make([]int, other.numCols)
	}

	for i, row := range m.els {
		for j := 0; j < other.numCols; j++ {
			col := other.col(j)
			var res int
			for k := range row {
				res += row[k] * col[k]
			}
			els[i][j] = res
			if res > 1 {
				els[i][j] = 1
			}
		}
	}

	return NewBooleanMatrix(els)
}

func (m *BooleanMatrix) Join(other *BooleanMatrix) (*BooleanMatrix, error) {
	if m.numRows != other.numRows || m.numCols != other.numCols {
		return nil, fmt.Errorf("matrices should be of the same dimensions")
	}

	els := make([][]int, m.numRows)
	for i := range els {
		els[i] = make([]int, other.numCols)
	}

	for i, row := range m.els {
		for j, el := range row {
			els[i][j] = el + other.els[i][j]
			if els[i][j] > 1 {
				els[i][j] = 1
			}
		}
	}

	return NewBooleanMatrix(els)
}

func (m *BooleanMatrix) Meet(other *BooleanMatrix) (*BooleanMatrix, error) {
	if m.numRows != other.numRows || m.numCols != other.numCols {
		return nil, fmt.Errorf("matrices should be of the same dimensions")
	}

	els := make([][]int, m.numRows)
	for i := range els {
		els[i] = make([]int, other.numCols)
	}

	for i, row := range m.els {
		for j, el := range row {
			els[i][j] = el * other.els[i][j]
		}
	}

	return NewBooleanMatrix(els)
}

// Power raises square matrix to power n
func (sm *BooleanMatrix) Power(n int) (*BooleanMatrix, error) {
	if n < 1 {
		return nil, fmt.Errorf("power should be greater that zero")
	}

	if sm.numRows != sm.numCols {
		return nil, fmt.Errorf("matrix should be square")
	}

	product := sm
	var err error

	for i := 2; i <= n; i++ {
		product, err = product.X(sm)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}
