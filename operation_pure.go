package matrix

/*
Perform an addition operation on two matrixes creating a new matrix with the results
*/
func (m Matrix[T]) Add(a *Matrix[T]) (*Matrix[T], error) {
	if !AreSameDimensions(&m, a) {
		return nil, ErrMustBeSameDimensions
	}

	result, err := NewEmptyMatrix[T](m.rows, m.columns)
	if err != nil {
		return nil, err
	}

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			newVal := m.reader.Read(i, j) + a.reader.Read(i, j)
			result.writer.Write(i, j, newVal)
		}
	}

	return result, nil
}

func (m Matrix[T]) Subtract(a *Matrix[T]) (*Matrix[T], error) {
	if !AreSameDimensions(&m, a) {
		return nil, ErrMustBeSameDimensions
	}

	result, err := NewEmptyMatrix[T](m.rows, m.columns)
	if err != nil {
		return nil, err
	}

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			newVal := m.reader.Read(i, j) - a.reader.Read(i, j)
			result.writer.Write(i, j, newVal)
		}
	}

	return result, nil
}

/*
Performs scalar multiplication on a matrix returning a new result matrix
*/
func (m Matrix[T]) ScalarMultiply(c T) (*Matrix[T], error) {
	result, err := NewEmptyMatrix[T](m.rows, m.columns)
	if err != nil {
		return nil, err
	}

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			newVal := c * m.reader.Read(i, j)
			result.writer.Write(i, j, newVal)
		}
	}

	return result, nil
}

func (m Matrix[T]) Multiply(a *Matrix[T]) (*Matrix[T], error) {
	// The number of columns in the first matrix must be equal to the number of rows in the second
	if m.columns != a.rows {
		return nil, ErrMultiplicationColumnRowMismatch
	}

	// The new matrix has the number of rows of the first and the number of columns of the second
	new, _ := NewEmptyMatrix[T](m.rows, a.columns)

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < a.columns; j++ {
			var sum T
			for k := uint(0); k < m.columns; k++ {
				sum += m.reader.Read(i, k) * a.reader.Read(k, j)
			}
			new.writer.Write(i, j, sum)
		}
	}

	return new, nil
}

/*
Search the given matrix for an element. Returns a list of Location with
position and values and boolean noting if it was found or not.
*/
func (m Matrix[T]) Search(element T) ([]Location[T], bool) {
	var found []Location[T]

	if m.HasIndex() {
		indexResult, ok := m.index[element]

		if ok {
			for _, v := range indexResult {
				found = append(found, Location[T]{
					position: v,
					value:    element,
				})
			}
			return found, true
		} else {
			return found, false
		}
	}

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			if m.reader.Read(i, j) == element {
				el := Location[T]{
					position: [2]uint{i, j},
					value:    m.reader.Read(i, j),
				}
				found = append(found, el)
			}
		}
	}

	if len(found) == 0 {
		return found, false
	}

	return found, true
}

func (m Matrix[T]) Transpose() (*Matrix[T], error) {
	new, err := NewEmptyMatrix[T](m.columns, m.rows)
	if err != nil {
		return nil, err
	}

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			new.writer.Write(j, i, m.reader.Read(i, j))
		}
	}

	return new, nil
}

func (m Matrix[T]) Flatten() []T {
	result := []T{}

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			result = append(result, m.reader.Read(i, j))
		}
	}

	return result
}

/*
Expands a slice into a new matrix of the given dimensions, will return an error if the
matrix does not have enough space to fit the slice. If there are more elements in the
new matrix than the input slice the remaining matrix elements will be zero filled.
*/
func ExpandSliceToMatrix[T Element](values []T, rows, columns uint) (*Matrix[T], error) {
	matrixSize := rows * columns
	inputSize := uint(len(values))

	if inputSize > matrixSize {
		return nil, ErrMatrixOverflow(matrixSize, inputSize)
	}

	new, err := NewEmptyMatrix[T](rows, columns)

	if err != nil {
		return nil, err
	}

	h := 0

	for i := uint(0); i < new.rows; i++ {
		for j := uint(0); j < new.columns; j++ {
			if h < len(values) {
				new.writer.Write(i, j, values[h])
				h++
			} else {
				break
			}
		}
	}

	return new, nil
}
