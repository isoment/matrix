package matrix

import (
	"errors"
	"fmt"
)

// Some errors don't need any additional formatted data
var (
	ErrRowColumSize         = errors.New("rows and columns must be greater than zero")
	ErrMustBeSameDimensions = errors.New("matrixes must be the same dimensions")
	ErrMatrixOutOfBounds    = errors.New("provided position exceeds matrix dimensions")
)

func ErrColumnCountMismatch(row int) error {
	return fmt.Errorf("column count mismatch in row %d", row)
}

func ErrRowCountMismatch(dataRows, rows uint) error {
	return fmt.Errorf("mismatch between data rows: %d and rows param: %d", dataRows, rows)
}

func ErrMatrixOverflow(matrixSize, inputSize uint) error {
	return fmt.Errorf("matrix has size: %d cannot fit input size: %d", matrixSize, inputSize)
}
