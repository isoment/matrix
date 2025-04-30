package matrix

/*
Represents an element in a matrix. Might consider https://pkg.go.dev/golang.org/x/exp/constraints
for use in the future.
*/
type Element interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type DefaultDataStore[T Element] struct {
	data [][]T
}

type DataReader[T Element] interface {
	Read(i, j uint) T
	Shape() (rows, columns uint)
	Validate() error
}

type DataWriter[T Element] interface {
	Write(i, j uint, value T)
}

// Read a value from a position in the matrix
func (d *DefaultDataStore[T]) Read(i, j uint) T {
	return d.data[i][j]
}

// Get the dimensions of the underlying matrix data store
func (d *DefaultDataStore[T]) Shape() (uint, uint) {
	if len(d.data) == 0 {
		return 0, 0
	}
	if len(d.data[0]) == 0 {
		return uint(len(d.data)), 0
	}
	return uint(len(d.data)), uint(len(d.data[0]))
}

// Ensure that the matrix has a valid structure
func (d *DefaultDataStore[T]) Validate() error {
	if len(d.data) == 0 || len(d.data[0]) == 0 {
		return ErrRowColumSize
	}

	err := verifyColumnCount(d.data)
	if err != nil {
		return err
	}

	return nil
}

// Write a value to a position in the matrix
func (d *DefaultDataStore[T]) Write(i, j uint, value T) {
	d.data[i][j] = value
}

// The index is an optional field to help speed lookup when needed
type Matrix[T Element] struct {
	rows    uint
	columns uint
	index   map[T][][2]uint
	reader  DataReader[T]
	writer  DataWriter[T]
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

func (m *Matrix[T]) HasIndex() bool {
	if m.index != nil {
		return true
	}
	return false
}

func (m *Matrix[T]) Size() uint {
	return m.rows * m.columns
}

/*
Create a new matrix specifying the size and data
*/
func NewMatrix[T Element](reader DataReader[T]) (*Matrix[T], error) {
	rows, columns := reader.Shape()

	err := reader.Validate()
	if err != nil {
		return nil, err
	}

	matrix := &Matrix[T]{
		rows:    rows,
		columns: columns,
		reader:  reader,
	}

	// Set the DataWriter on the matrix
	if w, ok := reader.(DataWriter[T]); ok {
		matrix.writer = w
	}

	return matrix, nil
}

func NewMatrixFromSlice[T Element](data [][]T) (*Matrix[T], error) {
	if len(data) == 0 || len(data[0]) == 0 {
		return nil, ErrRowColumSize
	}

	err := verifyColumnCount(data)
	if err != nil {
		return nil, err
	}

	store := &DefaultDataStore[T]{data: data}
	return NewMatrix(store)
}

/*
Create a new empty matrix with a given size
*/
func NewEmptyMatrix[T Element](rows, columns uint) (*Matrix[T], error) {
	if rows == 0 || columns == 0 {
		return nil, ErrRowColumSize
	}

	data := make([][]T, rows)
	for i := range data {
		data[i] = make([]T, columns)
	}

	store := &DefaultDataStore[T]{data: data}
	return NewMatrix(store)
}

/*
Build an index of the matrix for cases where quicker lookup might be desired. Adds
overhead for index storage and maintenance.
*/
func (m *Matrix[T]) Index() error {
	if m.HasIndex() {
		return ErrIndexExists
	}

	index := make(map[T][][2]uint)

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			element := m.reader.Read(i, j)

			v, ok := index[element]
			if ok {
				index[element] = append(v, [2]uint{i, j})
			} else {
				index[element] = [][2]uint{{i, j}}
			}
		}
	}

	m.index = index
	return nil
}
