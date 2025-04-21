package matrix

func (m *Matrix[T]) AddInPlace(a *Matrix[T]) (*Matrix[T], error) {
	if !AreSameDimensions(m, a) {
		return nil, ErrMustBeSameDimensions
	}

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			m.data[i][j] = m.data[i][j] + a.data[i][j]
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
			m.data[i][j] = m.data[i][j] * c
		}
	}
	return m
}
