package matrix

/*
Set a value at the given position. Row i and column j are 0 indexed
*/
func (m *Matrix[T]) Set(i, j uint, value T) (*Matrix[T], error) {
	if i >= m.rows || j >= m.columns {
		return nil, ErrMatrixOutOfBounds
	}

	m.writer.Write(i, j, value)
	return m, nil
}

func (m *Matrix[T]) AddInPlace(a *Matrix[T]) (*Matrix[T], error) {
	if !AreSameDimensions(m, a) {
		return nil, ErrMustBeSameDimensions
	}

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			newVal := m.reader.Read(i, j) + a.reader.Read(i, j)
			m.writer.Write(i, j, newVal)
		}
	}
	return m, nil
}

func (m *Matrix[T]) SubtractInPlace(a *Matrix[T]) (*Matrix[T], error) {
	if !AreSameDimensions(m, a) {
		return nil, ErrMustBeSameDimensions
	}

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			newVal := m.reader.Read(i, j) - a.reader.Read(i, j)
			m.writer.Write(i, j, newVal)
		}
	}

	return m, nil
}

/*
Performs the scalar multiplication operation but on the original matrix
*/
func (m *Matrix[T]) ScalarMultiplyInPlace(c T) *Matrix[T] {
	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			newVal := m.reader.Read(i, j) * c
			m.writer.Write(i, j, newVal)
		}
	}
	return m
}

func (m *Matrix[T]) HadamardProductInPlace(a *Matrix[T]) (*Matrix[T], error) {
	if !AreSameDimensions(m, a) {
		return nil, ErrMustBeSameDimensions
	}

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			product := m.reader.Read(i, j) * a.reader.Read(i, j)
			m.writer.Write(i, j, product)
		}
	}

	return m, nil
}
