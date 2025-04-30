package matrix

import (
	"errors"
	"fmt"
)

var (
	ErrRowColumSize                    = errors.New("rows and columns must be greater than zero")
	ErrMustBeSameDimensions            = errors.New("matrixes must be the same dimensions")
	ErrMatrixOutOfBounds               = errors.New("provided position exceeds matrix dimensions")
	ErrIndexExists                     = errors.New("matrix already has an index")
	ErrMultiplicationColumnRowMismatch = errors.New("param matrix row count must match receiver matrix column count")
)

func ErrColumnCountMismatch(row int) error {
	return fmt.Errorf("column count mismatch in row %d", row)
}

func ErrMatrixOverflow(matrixSize, inputSize uint) error {
	return fmt.Errorf("matrix has size: %d cannot fit input size: %d", matrixSize, inputSize)
}
