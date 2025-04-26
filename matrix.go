package matrix

/*
Represents an element in a matrix. Might consider https://pkg.go.dev/golang.org/x/exp/constraints
for use in the future.
*/
type Element interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// The index is an optional field to help speed up searching when needed
type Matrix[T Element] struct {
	rows    uint
	columns uint
	data    [][]T
	index   map[T][][2]uint
}

// The position field is the row, column coordinates 0 indexed
type Location[T Element] struct {
	position [2]uint
	value    T
}

func (m *Matrix[T]) Rows() uint {
	return m.rows
}

func (m *Matrix[T]) Columns() uint {
	return m.columns
}

func (m Matrix[T]) HasIndex() bool {
	if m.index != nil {
		return true
	}
	return false
}

func (m Matrix[T]) Size() uint {
	return m.rows * m.columns
}

/*
Create a new matrix specifying the size and data
*/
func NewMatrix[T Element](rows, columns uint, data [][]T) (*Matrix[T], error) {
	if rows == 0 || columns == 0 {
		return nil, ErrRowColumSize
	}
	if uint(len(data)) != rows {
		return nil, ErrRowCountMismatch(uint(len(data)), rows)
	}
	for i, row := range data {
		if uint(len(row)) != columns {
			return nil, ErrColumnCountMismatch(i)
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
		return nil, ErrRowColumSize
	}
	new := createEmptyMatrix[T](rows, columns)
	return &new, nil
}

func (m *Matrix[T]) Index() {

}
