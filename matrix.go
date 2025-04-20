package matrix

import (
	"errors"
	"fmt"
)

/*
Represents an element in a matrix. Might consider https://pkg.go.dev/golang.org/x/exp/constraints
for use in the future.
*/
type Element interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Matrix[T Element] struct {
	rows    uint
	columns uint
	data    [][]T
}

func (m *Matrix[T]) Rows() uint {
	return m.rows
}

func (m *Matrix[T]) Columns() uint {
	return m.columns
}

/*
Create a new matrix specifying the size and data
*/
func NewMatrix[T Element](rows, columns uint, data [][]T) (*Matrix[T], error) {
	if rows == 0 || columns == 0 {
		return nil, errors.New("matrix rows and columns must be greater than zero")
	}
	if uint(len(data)) != rows {
		return nil, errors.New("row count mismatch")
	}
	for i, row := range data {
		if uint(len(row)) != columns {
			return nil, fmt.Errorf("column count mismatch in row %d", i)
		}
	}

	return &Matrix[T]{
		rows:    rows,
		columns: columns,
		data:    data,
	}, nil
}

/*
Create a new empty matrix with a given size
*/
func NewEmptyMatrix[T Element](rows, columns uint) (*Matrix[T], error) {
	if rows == 0 || columns == 0 {
		return nil, errors.New("matrix rows and columns must be greater than zero")
	}
	new := createEmptyMatrix[T](rows, columns)
	return &new, nil
}
