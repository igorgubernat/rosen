package matrix

import (
	"fmt"
)

type SquareMatrix struct {
	Matrix
}

func NewSquareMatrix(els [][]float64) (*SquareMatrix, error) {
	m, err := NewMatrix(els)
	if err != nil {
		return nil, err
	}

	if m.numRows != m.numCols {
		return nil, fmt.Errorf("matrix %v is not square", m)
	}

	return &SquareMatrix{
		Matrix: *m,
	}, nil
}

// X multiplies two square matrices
func (sm *SquareMatrix) X(other *SquareMatrix) (*SquareMatrix, error) {
	m, err := sm.Matrix.X(&other.Matrix)
	if err != nil {
		return nil, err
	}

	return &SquareMatrix{
		Matrix: *m,
	}, nil
}

// Power raises square matrix to power n
func (sm *SquareMatrix) Power(n int) (*SquareMatrix, error) {
	if n < 1 {
		return nil, fmt.Errorf("power should be greater that zero")
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

// IsSymmetric returns true if matrix is symmetric
func (sm *SquareMatrix) IsSymmetric() bool {
	for i, row := range sm.Matrix.els {
		for j, el := range row {
			if el != sm.Matrix.els[j][i] {
				return false
			}
		}
	}

	return true
}
