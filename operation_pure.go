package matrix

/*
Perform an addition operation on two matrixes creating a new matrix with the results
*/
func (m Matrix[T]) Add(a *Matrix[T]) *Matrix[T] {
	rows := uint(len(m.data))
	columns := uint(len(m.data[0]))
	result := createEmptyMatrix[T](rows, columns)

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			result.data[i][j] = m.data[i][j] + a.data[i][j]
		}
	}

	return &result
}

/*
Performs scalar multiplication on a matrix returning a new result matrix
*/
func (m Matrix[T]) ScalarMultiply(c T) *Matrix[T] {
	result := createEmptyMatrix[T](m.rows, m.columns)

	for i := uint(0); i < m.rows; i++ {
		for j := uint(0); j < m.columns; j++ {
			result.data[i][j] = c * m.data[i][j]
		}
	}

	return &result
}
